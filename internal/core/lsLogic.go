package core

import (
	"fmt"

	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils"
	"github.com/silentFellow/cred-store/internal/utils/paths"
)

func LsLogic(
	cmdType string,
	args []string,
) {
	var basePath string
	if cmdType == "pass" {
		basePath = config.Constants.PassPath
	} else {
		basePath = config.Constants.EnvPath
	}

	if len(args) < 1 {
		if err := utils.PrintTree(basePath, "", true); err != nil {
			fmt.Println("Failed to list the files and directories, ", err)
		}
		return
	}

	path := args[0]
	fullPath := paths.BuildPath(basePath, path)
	if err := utils.PrintTree(fullPath, "", true); err != nil {
		fmt.Println("Failed to list the files and directories, ", err)
	}
}
