package gpgcrypt

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/silentFellow/cred/config"
	"github.com/silentFellow/cred/internal/utils"
	"github.com/silentFellow/cred/internal/utils/copy"
	"github.com/silentFellow/cred/internal/utils/paths"
)

func GenerateKey(uname, email string) error {
	identity := fmt.Sprintf("%v <%v>", uname, email)

	cmd := utils.SetCmd(
		utils.CmdConfig{},
		"gpg",
		"--quick-generate-key",
		identity,
		"rsa4096",
		"default",
		"never",
	)

	if output, err := cmd.CombinedOutput(); err != nil &&
		!strings.Contains(string(output), "already exists") {
		return err
	}

	return nil
}

func GetKeyFpr(uname string) (string, error) {
	cmd := utils.SetCmd(utils.CmdConfig{}, "gpg", "--list-keys", "--with-colons", uname)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(string(output), "\n") {
		fields := strings.Split(line, ":")
		if fields[0] == "fpr" {
			if len(fields) < 11 {
				return "", errors.New("Invalid key format")
			}
			return fields[9], nil
		}
	}

	return "", nil
}

func AddSubKey(keyId string) error {
	cmd := utils.SetCmd(
		utils.CmdConfig{IsStdin: true},
		"gpg",
		"--command-fd",
		"0",
		"--batch",
		"--yes",
		"--edit-key",
		keyId,
	)

	// "addkey" -> Select subkey type (6 = RSA) -> Key size (4096) -> Usage (E = Encrypt) -> Save
	cmd.Stdin = strings.NewReader("addkey\n6\n\n\ny\ny\nsave\n")

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func ModifyTrust(keyId string) error {
	cmd := utils.SetCmd(
		utils.CmdConfig{IsStdin: true},
		"gpg",
		"--command-fd",
		"0",
		"--batch",
		"--yes",
		"--edit-key",
		keyId,
	)

	// "addkey" -> Select subkey type (6 = RSA) -> Key size (4096) -> Usage (E = Encrypt) -> Save
	cmd.Stdin = strings.NewReader("trust\n5\ny\nsave\n")

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func ExportKeys(uname string) error {
	tempDir, err := os.MkdirTemp("", "cred-store-keys-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir) // Cleanup temporary directory

	publicKeyPath := paths.BuildPath(tempDir, "public_key.asc")
	publicKeyCmd := utils.SetCmd(
		utils.CmdConfig{IsStdout: true},
		"gpg",
		"--armor",
		"--export",
		uname,
	)
	publicKeyFile, err := os.Create(publicKeyPath)
	if err != nil {
		return err
	}
	defer publicKeyFile.Close()

	publicKeyCmd.Stdout = publicKeyFile
	if err := publicKeyCmd.Run(); err != nil {
		return err
	}

	privateKeyPath := paths.BuildPath(tempDir, "private_key.asc")

	privateKeyCmd := utils.SetCmd(
		utils.CmdConfig{IsStdout: true},
		"gpg",
		"--armor",
		"--export-secret-key",
		uname,
	)
	privateKeyFile, err := os.Create(privateKeyPath)
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()

	privateKeyCmd.Stdout = privateKeyFile
	if err := privateKeyCmd.Run(); err != nil {
		return err
	}

	// Generate a usage file for importing the keys
	usageFilePath := paths.BuildPath(tempDir, "import_instructions.txt")
	usageFileContent := `To import the keys on a new device:

1. Import the public key:
   gpg --import /path/to/public_key.asc

2. Import the private key:
   gpg --import /path/to/private_key.asc

3. Optionally, set the trust level:
   gpg --edit-key <key_id> trust
   (Set trust to ultimate for full trust.)
	 ("save" to save and exit)

4. Verify the keys are imported:
   gpg --list-keys
`

	usageFile, err := os.Create(usageFilePath)
	if err != nil {
		return err
	}
	defer usageFile.Close()

	if _, err := usageFile.WriteString(usageFileContent); err != nil {
		return err
	}

	downloadPath := paths.BuildPath(config.Constants.Download, "cred-store-keys")
	if err := fscopy.Copy(tempDir, downloadPath); err != nil {
		return err
	}

	return nil
}
