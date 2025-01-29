package core

import (
	"fmt"
	"os"

	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils/paths"
)

func MkdirLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v mkdir <directory>", cmdType)

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

	for _, path := range args {
		fullPath := paths.BuildPath(basePath, path)

		if err := os.MkdirAll(fullPath, 0777); err != nil {
			fmt.Printf("failed to create directory %v: %v\n", fullPath, err)
			continue
		}

		fmt.Printf("%v created successfully\n", path)
	}
}
