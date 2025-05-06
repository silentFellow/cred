package core

import (
	"fmt"

	"github.com/silentFellow/cred/config"
	"github.com/silentFellow/cred/internal/utils/copy"
	"github.com/silentFellow/cred/internal/utils/paths"
)

func CpLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v cp <src> <dest>", cmdType)

	var basePath string
	if cmdType == "pass" {
		basePath = config.Constants.PassPath
	} else if cmdType == "env" {
		basePath = config.Constants.EnvPath
	} else if cmdType == "ssh" {
		basePath = config.Constants.SshPath
	}

	if len(args) < 2 {
		fmt.Println("invalid usage, expected: ", usage)
		return
	}

	n := len(args)
	sources := args[:n-1]
	destination := args[n-1]
	destinationPath := paths.BuildPath(basePath, args[n-1])

	for _, src := range sources {
		srcPath := paths.BuildPath(basePath, src)
		if err := fscopy.Copy(srcPath, destinationPath); err != nil {
			fmt.Printf("copying %v failed: %v\n", src, err)
			continue
		}

		fmt.Printf("file copied from %v to %v\n", src, destination)
	}
}
