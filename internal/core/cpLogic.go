package core

import (
	"fmt"

	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils/copy"
	"github.com/silentFellow/cred-store/internal/utils/paths"
)

func CpLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v cp <src> <dest>", cmdType)

	var basePath string
	if cmdType == "pass" {
		basePath = config.Constants.PassPath
	} else {
		basePath = config.Constants.EnvPath
	}

	if len(args) < 2 {
		fmt.Printf("Invalid usage: %v\n", usage)
		return
	}

	n := len(args)
	sources := args[:n-1]
	destination := paths.BuildPath(basePath, args[n-1])

	for _, src := range sources {
		srcPath := paths.BuildPath(basePath, src)
		if err := fscopy.Copy(srcPath, destination); err != nil {
			fmt.Printf("Error copying %v: %v\n", srcPath, err)
			continue
		}

		fmt.Printf("Succesfully copied from %v to %v\n", srcPath, destination)
	}
}
