package sshUtils

import (
	sshCrypto "golang.org/x/crypto/ssh"

	"github.com/silentFellow/cred/internal/utils"
)

// function to check if ssh is installed
func CheckSshExists() bool {
	cmd := utils.SetCmd(utils.CmdConfig{}, "ssh", "-V")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// function to validate ssh public key
func ValidatePublicKey(key string) bool {
	_, _, _, _, err := sshCrypto.ParseAuthorizedKey([]byte(key))
	return err == nil
}

// function to validate ssh private key
func ValidatePrivateKey(key string) bool {
	_, err := sshCrypto.ParsePrivateKey([]byte(key))
	return err == nil
}

// function to validate ssh key
func ValidateKey(key string, keyType string) bool {
	if keyType != "public" && keyType != "private" {
		return false
	}

	if keyType == "public" {
		return ValidatePublicKey(key)
	} else {
		return ValidatePrivateKey(key)
	}
}
