package core

import (
	"fmt"
	"path/filepath"

	"github.com/silentFellow/cred/config"
	"github.com/silentFellow/cred/internal/utils"
	"github.com/silentFellow/cred/internal/utils/paths"
)

func LsLogic(
	cmdType string,
	args []string,
) {
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
		if err := utils.PrintTree(basePath, "", true); err != nil {
			fmt.Printf(
				"listing files and directories in %v failed: %v\n",
				filepath.Base(basePath),
				err,
			)
		}
		return
	}

	path := args[0]
	fullPath := paths.BuildPath(basePath, path)
	if err := utils.PrintTree(fullPath, "", true); err != nil {
		fmt.Printf("listing files and directories in %v failed: %v\n", path, err)
	}
}
