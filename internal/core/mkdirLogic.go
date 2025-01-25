package core

import (
	"fmt"
	"os"

	"github.com/silentFellow/cred-store/config"
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
		fmt.Printf("Invalid usage: %v\n", usage)
		return
	}

	for _, path := range args {
		fullPath := fmt.Sprintf("%v/%v", basePath, path)

		if err := os.MkdirAll(fullPath, 0777); err != nil {
			fmt.Printf("Error created %v: %v\n", fullPath, err)
			continue
		}

		fmt.Printf("Succesfully created %v\n", fullPath)
	}
}
