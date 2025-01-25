package core

import (
	"fmt"
	"os"

	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils"
)

func RmLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v rm <filepath>", cmdType)

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

		if !utils.CheckPathExists(fullPath) {
			fmt.Printf("%v not found\n", fullPath)
			continue
		}

		if err := os.RemoveAll(fullPath); err != nil {
			fmt.Printf("Error removing %v: %v\n", fullPath, err)
			continue
		}

		fmt.Printf("Succesfully deleted %v\n", fullPath)
	}
}
