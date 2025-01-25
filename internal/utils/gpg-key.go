package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/silentFellow/cred-store/config"
)

func CheckKeyExists() bool {
  gpgPath := fmt.Sprintf("%v/.gpg-id", config.Constants.StorePath)

  if _, err := os.Stat(gpgPath); err != nil {
    return false
  }

  return true
}

func CheckKeyValidity(keyId string) bool {
	cmd := exec.Command("gpg", "--list-keys", keyId)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	if strings.Contains(string(output), keyId) {
		return true
	}

	return false
}
