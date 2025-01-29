package gpgcrypt

import (
	"bytes"
	"errors"

	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils"
)

func Decrypt(filePath string) (string, error) {
	key := config.Constants.GpgKey

	if key == "" {
		return "", errors.New("Invalid GPG key")
	}

	cmd := utils.SetCmd("", utils.CmdIOConfig{}, "gpg", "--decrypt", filePath)

	var outBuffer bytes.Buffer
	cmd.Stdout = &outBuffer

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return outBuffer.String(), nil
}
