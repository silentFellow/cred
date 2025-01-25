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

	path := args[0]
	fullPath := fmt.Sprintf("%v/%v", basePath, path)

	if !utils.CheckPathExists(fullPath) {
		fmt.Println("Path not found")
		return
	}

  if err := os.RemoveAll(fullPath); err != nil {
    fmt.Println("Error removing path")
    return
  }

	fmt.Println("Succesfully deleted path")
}
