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
	switch cmdType {
	case "pass":
		basePath = config.Constants.PassPath
	case "env":
		basePath = config.Constants.EnvPath
	case "ssh":
		basePath = config.Constants.SshPath
	}

	if len(args) < 2 {
		fmt.Println("invalid usage, expected: ", usage)
		return
	}

	n := len(args)
	sources := args[:n-1]
	destination := args[n-1]
	destinationPath := paths.BuildPath(basePath, args[n-1])

	destinationInfo, err := os.Stat(destinationPath)
	if err != nil && !os.IsNotExist(err) { // only errors not related to being not found
		fmt.Printf("failed to check destination %v: %v\n", destination, err)
		return
	}

	if destinationInfo != nil && destinationInfo.IsDir() { // if directory
		for _, src := range sources {
			srcPath := paths.BuildPath(basePath, src)
			destPath := filepath.Join(destinationPath, filepath.Base(src)) // append
			move(srcPath, destPath, basePath)
		}
	} else { // if normal file
		for _, src := range sources {
			srcPath := paths.BuildPath(basePath, src)
			move(srcPath, destinationPath, basePath)
		}
	}
}

func move(source, destination, basePath string) {
	// for better success/error messages
	soruceRelativePath, srcRelErr := filepath.Rel(basePath, source)
	destRelativePath, destRelErr := filepath.Rel(basePath, destination)

	if err := os.Rename(source, destination); err != nil {
		if srcRelErr != nil || destRelErr != nil {
			fmt.Println("failed to move file: ", err)
		}
		fmt.Printf("failed to move %v to %v: %v\n", soruceRelativePath, destRelativePath, err)
		return
	}

	if srcRelErr != nil || destRelErr != nil {
		fmt.Println("successfully moved file")
		return
	}
	fmt.Printf("successfully moved from %v to %v\n", soruceRelativePath, destRelativePath)
}
