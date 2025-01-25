package gpgcrypt

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils/copy"
)

func GenerateKey(uname, email string) error {
	identity := fmt.Sprintf("%v <%v>", uname, email)
	cmd := exec.Command("gpg", "--quick-generate-key", identity, "rsa4096", "default", "never")

	output, err := cmd.CombinedOutput()
	formattedOutput := strings.ToLower(string(output))
	if err != nil &&
		!strings.Contains(
			formattedOutput,
			"key already exists",
		) { // TODO: key already exists throws an error
		return err
	}

	return nil
}

func GetKeyFpr(uname string) (string, error) {
	cmd := exec.Command("gpg", "--list-keys", "--with-colons", uname)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(string(output), "\n") {
		fields := strings.Split(line, ":")
		if fields[0] == "fpr" {
			if len(fields) < 11 {
				return "", fmt.Errorf("Invalid key format")
			}
			return fields[9], nil
		}
	}

	return "", nil
}

func AddSubKey(keyId string) error {
	cmd := exec.Command("gpg", "--command-fd", "0", "--batch", "--yes", "--edit-key", keyId)

	// "addkey" -> Select subkey type (6 = RSA) -> Key size (4096) -> Usage (E = Encrypt) -> Save
	cmd.Stdin = strings.NewReader("addkey\n6\n\n\ny\ny\nsave\n")

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func ModifyTrust(keyId string) error {
	cmd := exec.Command("gpg", "--command-fd", "0", "--batch", "--yes", "--edit-key", keyId)

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
		return fmt.Errorf("Failed to create keys-temp directory: %w", err)
	}

	publicKeyPath := fmt.Sprintf("%v/public_key.asc", tempDir)
	publicKeyCmd := exec.Command("gpg", "--armor", "--export", uname)
	publicKeyFile, err := os.Create(publicKeyPath)
	if err != nil {
		return fmt.Errorf("Failed to create public key file: %w", err)
	}
	defer publicKeyFile.Close()

	publicKeyCmd.Stdout = publicKeyFile
	if err := publicKeyCmd.Run(); err != nil {
		return fmt.Errorf("Failed to export public key: %w", err)
	}

	privateKeyPath := fmt.Sprintf("%v/private_key.asc", tempDir)
	privateKeyCmd := exec.Command("gpg", "--armor", "--export-secret-keys", uname)
	privateKeyFile, err := os.Create(privateKeyPath)
	if err != nil {
		return fmt.Errorf("Failed to create private key file: %w", err)
	}
	defer privateKeyFile.Close()

	privateKeyCmd.Stdout = privateKeyFile
	if err := privateKeyCmd.Run(); err != nil {
		return fmt.Errorf("Failed to export private key: %w", err)
	}

	// Generate a usage file for importing the keys
	usageFilePath := fmt.Sprintf("%v/import_instructions.txt", tempDir)
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
		return fmt.Errorf("Failed to create import_usage file: %w", err)
	}
	defer usageFile.Close()

	if _, err := usageFile.WriteString(usageFileContent); err != nil {
		return fmt.Errorf("Failed to write content to import_usage file: %w", err)
	}

	downloadPath := fmt.Sprintf("%v/cred-store-keys", config.Constants.Download)
	if err := fscopy.Copy(tempDir, downloadPath); err != nil {
		return fmt.Errorf("Failed to move keys from temp to download directory: %w", err)
	}

	return nil
}
