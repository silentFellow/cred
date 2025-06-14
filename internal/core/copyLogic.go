package core

import (
	"fmt"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils"
	"github.com/silentFellow/cred/internal/utils/paths"
)

func CopyLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v copy <filename>", cmdType)

	var basePath string
	if cmdType == "pass" {
		basePath = config.Constants.PassPath
	} else if cmdType == "env" {
		basePath = config.Constants.EnvPath
	} else if cmdType == "ssh" {
		basePath = config.Constants.SshPath
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
		fmt.Println("invalid file format, only file is allowed")
		return
	}

	decrypted, err := gpgcrypt.Decrypt(fullPath)
	if err != nil {
		fmt.Printf("decrypting file %v failed: %v\n", path, err)
		return
	}

	copyOnlyFirst := false
	if cmdType == "pass" {
		copyOnlyFirst = true
	}
	if err := utils.CopyToClipboard(decrypted, copyOnlyFirst); err != nil {
		fmt.Println("copying to clipboard failed: ", err)
		return
	}

	fmt.Println("successfully copies to clipboard")
}
