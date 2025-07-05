package core

import (
	"fmt"
	"os"

	"github.com/silentFellow/cred/config"
	"github.com/silentFellow/cred/internal/utils/paths"
)

func MkdirLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v mkdir <directory>", cmdType)

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

	for _, path := range args {
		fullPath := paths.BuildPath(basePath, path)

		if err := os.MkdirAll(fullPath, 0777); err != nil {
			fmt.Printf("failed to create directory %v: %v\n", path, err)
			continue
		}

		fmt.Printf("%v created successfully\n", path)
	}
}
