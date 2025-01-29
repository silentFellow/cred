package gpgcrypt

import (
	"os"
)

func Recrypt(path, oldKey, newKey string) error {
	originalContent, err := Decrypt(path)
	if err != nil {
		return err
	}

	encryptedContent, err := Encrypt(originalContent, newKey)
	if err != nil {
		return err
	}

	targetPath, err := os.Create(path)
	if err != nil {
		return err
	}
	defer targetPath.Close()

	if _, err := targetPath.WriteString(encryptedContent); err != nil {
		return err
	}

	return nil
}
