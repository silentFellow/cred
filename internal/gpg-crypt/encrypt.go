package gpgcrypt

import (
	"bytes"
	"errors"

	"github.com/silentFellow/cred/internal/utils"
)

func Encrypt(v string, key string) (string, error) {
	if key == "" {
		return "", errors.New("Invalid GPG key")
	}

	cmd := utils.SetCmd("", utils.CmdIOConfig{}, "gpg", "--armor", "--encrypt", "--recipient", key)

	// create a pipeline for input
	inPipe, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	// create a buffer to store output
	var outBuf bytes.Buffer
	cmd.Stdout = &outBuf

	// start the encryption
	if err := cmd.Start(); err != nil {
    inPipe.Close()
		return "", err
	}

	// write the content to GPG stdin
	if _, err := inPipe.Write([]byte(v)); err != nil {
    inPipe.Close()
		return "", err
	}
  inPipe.Close()

	// wait for GPG to finish the process
	if err := cmd.Wait(); err != nil {
		return "", err
	}

	return outBuf.String(), nil
}
