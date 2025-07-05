package gpgcrypt

import (
	"bytes"
	"errors"
	"strings"

	"github.com/silentFellow/cred/config"
	"github.com/silentFellow/cred/internal/utils"
)

func Decrypt(filePath string) (string, error) {
	key := config.Constants.GpgKey

	if key == "" {
		return "", errors.New("Invalid GPG key")
	}

	cmd := utils.SetCmd(utils.CmdConfig{}, "gpg", "--decrypt", filePath)

	var outBuffer bytes.Buffer
	var errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	if err := cmd.Run(); err != nil {
		errMsg := strings.ToLower(errBuffer.String())
		if strings.Contains(errMsg, "no passphrase") || strings.Contains(errMsg, "bad passphrase") {
			return "", errors.New("wrong passphrase")
		}

		return "", err
	}

	return outBuffer.String(), nil
}
