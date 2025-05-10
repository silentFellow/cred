package core

import (
	"fmt"
	"os"

	"github.com/silentFellow/cred/config"
	"github.com/silentFellow/cred/internal/utils/paths"
)

func RmLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v rm <filepath>", cmdType)

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

	for _, path := range args {
		fullPath := paths.BuildPath(basePath, path)

		if !paths.CheckPathExists(fullPath) {
			fmt.Printf("%v not found\n", path)
			continue
		}

		if err := os.RemoveAll(fullPath); err != nil {
			fmt.Printf("failed to remove %v: %v\n", path, err)
			continue
		}

		fmt.Printf("%v deleted successfully\n", path)
	}
}
