package core

import (
	"fmt"

	"github.com/silentFellow/cred-store/config"
	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils/paths"
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
		fmt.Println("invalid usage, expected: ", usage)
		return
	}
	path := args[0]
	fullPath := paths.BuildPath(basePath, path)

	if !paths.CheckPathExists(fullPath) {
		fmt.Printf("%v not found\n", path)
		return
	}

	if paths.GetPathType(fullPath) != "file" {
		fmt.Println("invalid file format, only files are allowed")
		return
	}

	decryped, err := gpgcrypt.Decrypt(fullPath)
	if err != nil {
		fmt.Println("failed to decrypt file: ", err)
		return
	}

	if cmdType == "pass" {
		fmt.Printf("%v\n", decryped)
	} else {
		fmt.Print(decryped)
	}
}
