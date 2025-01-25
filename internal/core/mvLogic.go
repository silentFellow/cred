package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/silentFellow/cred-store/config"
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
		fmt.Printf("Invalid usage: %v\n", usage)
		return
	}

	n := len(args)
	sources := args[:n-1]
	destination := fmt.Sprintf("%v/%v", basePath, args[n-1])

	destinationInfo, err := os.Stat(destination)
	if err != nil && !os.IsNotExist(err) { // only errors not related to being not found
		fmt.Println("Failed to check destination, ", err)
		return
	}

	if destinationInfo != nil && destinationInfo.IsDir() { // if directory
		for _, src := range sources {
			srcPath := fmt.Sprintf("%v/%v", basePath, src)
			destPath := filepath.Join(destination, filepath.Base(src)) // append
			move(srcPath, destPath)
		}
	} else { // if normal file
		for _, src := range sources {
			srcPath := fmt.Sprintf("%v/%v", basePath, src)
			move(srcPath, destination)
		}
	}
}

func move(source, destination string) {
	if err := os.Rename(source, destination); err != nil {
		fmt.Printf("Error moving %v: %v\n", source, err)
		return
	}

	fmt.Printf("Succesfully moved from %v to %v\n", source, destination)
}
