package gpgcrypt

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"github.com/silentFellow/cred-store/config"
)

func Encrypt(v string) (string, error) {
	key := config.Constants.GpgKey
	if key == "" {
		fmt.Println("Invalid GPG key")
		return "", errors.New("Invalid GPG key")
	}

	cmd := exec.Command("gpg", "--armor", "--encrypt", "--recipient", key)

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
