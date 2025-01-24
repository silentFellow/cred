package pass

import (
	"fmt"

	"github.com/atotto/clipboard"

	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils"
)

func addToPath(path string, content string, copy bool) error {
	file, err := utils.CreatePath(path)
	defer file.Close()

	if err != nil {
		return err
	}

	encrypted, err := gpgcrypt.Encrypt(content)
	if err != nil {
		return err
	}

	if _, err := file.WriteString(encrypted); err != nil {
		return err
	}

	if copy {
		if err := clipboard.WriteAll(content); err != nil {
			return err
		}
	}

	if copy {
		fmt.Println("Password inserted successfully, copied to clipboard")
	} else {
		fmt.Println("Password inserted successfully")
	}
	return nil
}
