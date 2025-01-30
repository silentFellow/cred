package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/silentFellow/cred/config"
	"github.com/silentFellow/cred/internal/utils/paths"
)

func MvLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v mv <src> <dest>", cmdType)

	var basePath string
	if cmdType == "pass" {
		basePath = config.Constants.PassPath
	} else {
		basePath = config.Constants.EnvPath
	}

	if len(args) < 2 {
		fmt.Println("invalid usage, expected: ", usage)
		return
	}

	n := len(args)
	sources := args[:n-1]
	destination := paths.BuildPath(basePath, args[n-1])

	destinationInfo, err := os.Stat(destination)
	if err != nil && !os.IsNotExist(err) { // only errors not related to being not found
		fmt.Printf("failed to check destination %v: %v\n", destination, err)
		return
	}

	if destinationInfo != nil && destinationInfo.IsDir() { // if directory
		for _, src := range sources {
			srcPath := paths.BuildPath(basePath, src)
			destPath := filepath.Join(destination, filepath.Base(src)) // append
			move(srcPath, destPath)
		}
	} else { // if normal file
		for _, src := range sources {
			srcPath := paths.BuildPath(basePath, src)
			move(srcPath, destination)
		}
	}
}

func move(source, destination string) {
	if err := os.Rename(source, destination); err != nil {
		fmt.Printf("failed to move %v to %v: %v\n", source, destination, err)
		return
	}

	fmt.Printf("successfully moved from %v to %v\n", source, destination)
}
