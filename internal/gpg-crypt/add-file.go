package gpgcrypt

import (
	"github.com/silentFellow/cred/config"
	"github.com/silentFellow/cred/internal/utils"
	"github.com/silentFellow/cred/internal/utils/paths"
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
		if err := utils.CopyToClipboard(content, false); err != nil {
			return err
		}
	}

	return nil
}
