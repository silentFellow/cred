package core

import (
	"fmt"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils/paths"
)

func ShowLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v show <filename>", cmdType)

	var basePath string
	switch cmdType {
	case "pass":
		basePath = config.Constants.PassPath
	case "env":
		basePath = config.Constants.EnvPath
	case "ssh":
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
		fmt.Println("invalid file format, only files are allowed")
		return
	}

	decrypted, err := gpgcrypt.Decrypt(fullPath)
	if err != nil {
		fmt.Println("failed to decrypt file: ", err)
		return
	}

	// add new line, if the decrypted content does not end with a new line
	if decrypted[len(decrypted)-1] != '\n' {
		fmt.Printf("%v\n", decrypted)
	} else {
		fmt.Print(decrypted)
	}
}
