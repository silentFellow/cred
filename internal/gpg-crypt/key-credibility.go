package gpgcrypt

import (
	"fmt"
	"os"
	"strings"

	"github.com/silentFellow/cred/config"
	"github.com/silentFellow/cred/internal/utils"
)

func CheckKeyExists() bool {
	gpgPath := fmt.Sprintf("%v/.gpg-id", config.Constants.StorePath)

	if _, err := os.Stat(gpgPath); err != nil {
		return false
	}

	return true
}

func CheckKeyValidity(keyId string) bool {
	cmd := utils.SetCmd(utils.CmdConfig{}, "gpg", "--list-keys", keyId)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	if strings.Contains(string(output), keyId) {
		return true
	}

	return false
}
