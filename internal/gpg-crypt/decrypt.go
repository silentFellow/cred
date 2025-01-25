package gpgcrypt

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"github.com/silentFellow/cred-store/config"
)

func Decrypt(filePath string) (string, error) {
	key := config.Constants.GpgKey

	if key == "" {
		fmt.Println("Invalid GPG key")
		return "", errors.New("Invalid GPG key")
	}

  cmd := exec.Command("gpg", "--decrypt", filePath)

  var outBuffer bytes.Buffer
  cmd.Stdout = &outBuffer

  if err := cmd.Run(); err != nil {
		fmt.Println(err)
    return "", err
  }

  return outBuffer.String(), nil
}
