package gpgcrypt

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils"
)

func Decrypt(filePath string) (string, error) {
	key := config.Constants.GpgKey

	if key == "" {
		fmt.Println("Invalid GPG key")
		return "", errors.New("Invalid GPG key")
	}

	cmd := utils.SetCmd("", utils.CmdIOConfig{}, "gpg", "--decrypt", filePath)

	var outBuffer bytes.Buffer
	cmd.Stdout = &outBuffer

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return "", err
	}

	return outBuffer.String(), nil
}
