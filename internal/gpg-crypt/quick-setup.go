package gpgcrypt

import (
	"fmt"
	"os/exec"
	"strings"
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
