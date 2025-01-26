package core

import (
	"fmt"

	"github.com/silentFellow/cred-store/config"
	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils"
	"github.com/silentFellow/cred-store/internal/utils/paths"
)

func CopyLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v copy <filename>", cmdType)

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

	if !paths.CheckPathExists(fullPath) {
		fmt.Printf("%v not found\n", path)
		return
	}

	if paths.GetPathType(fullPath) != "file" {
		fmt.Println("Invalid file format, only file is allowed")
		return
	}

	decryped, err := gpgcrypt.Decrypt(fullPath)
	if err != nil {
		fmt.Println("Error decrypting file")
		return
	}

	copyOnlyFirst := false
	if cmdType == "pass" {
		copyOnlyFirst = true
	}
	if err := utils.CopyToClipboard(decryped, copyOnlyFirst); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully copies to clipboard")
}
