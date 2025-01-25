package gpgcrypt

import (
	"fmt"
	"os"
)

func Recrypt(path, oldKey, newKey string) error {
	originalContent, err := Decrypt(path)
	if err != nil {
		return fmt.Errorf("Failed to decrypt the file: %v: ", err)
	}

	encryptedContent, err := Encrypt(originalContent, newKey)
	if err != nil {
		return fmt.Errorf("Failed to encrypt the file: %v: ", err)
	}

	targetPath, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create the target file %v: ", err)
	}
	defer targetPath.Close()

	if _, err := targetPath.WriteString(encryptedContent); err != nil {
		return fmt.Errorf("Failed to write to the target file %v: ", err)
	}

	return nil
}
