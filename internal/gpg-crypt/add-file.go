package gpgcrypt

import (
	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils"
	"github.com/silentFellow/cred-store/internal/utils/paths"
)

func AddFile(path string, content string, copy bool) error {
	file, err := paths.CreatePath(path)
	defer file.Close()

	if err != nil {
		return err
	}

	encrypted, err := Encrypt(content, config.Constants.GpgKey)
	if err != nil {
		return err
	}

	if _, err := file.WriteString(encrypted); err != nil {
		return err
	}

	if copy {
		// true since it only applicable for generate password
		if err := utils.CopyToClipboard(content, true); err != nil {
			return err
		}
	}

	return nil
}

