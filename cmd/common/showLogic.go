package common

import (
	"fmt"

	"github.com/silentFellow/cred-store/config"
	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils"
)

func ShowLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v show <filename>", cmdType)

	var basePath string
	if cmdType == "pass" {
		basePath = config.Constants.PassPath
	} else {
		basePath = config.Constants.EnvPath
	}

	if len(args) < 1 {
		fmt.Printf("Invalid usage: %v\n", usage)
		return
	}
	path := args[0]
	fullPath := fmt.Sprintf("%v/%v", basePath, path)

	if !utils.CheckPathExists(fullPath) {
		fmt.Println("Path not found")
		return
	}

	if utils.GetPathType(fullPath) != "file" {
		fmt.Println("Invalid file format, only file is allowed")
		return
	}

	decryped, err := gpgcrypt.Decrypt(fullPath)
	if err != nil {
		fmt.Println("Error decrypting file")
		return
	}

	fmt.Println(decryped)
}
