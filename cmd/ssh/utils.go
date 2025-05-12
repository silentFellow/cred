package ssh

import (
	"fmt"
	"os"

	"github.com/silentFellow/cred/internal/utils/paths"
)

// function to remove created paths
func removeCreated(path string) {
	if err := os.RemoveAll(path); err != nil {
		fmt.Println("failed to remove the created files: ", err)
	}
}

// function to prepare the key folder
func prepareKeyFolder(path string) bool {
	if paths.CheckPathExists(path) {
		var choice string
		fmt.Print("The ssh key already exists. Do you want to overwrite it? (y/n): ")
		fmt.Scanln(&choice)

		if choice != "y" && choice != "Y" {
			return false
		}

		if err := os.RemoveAll(path); err != nil {
			fmt.Println("failed to remove the file: ", err)
			return false
		}
	}

	if err := os.MkdirAll(path, 0700); err != nil {
		fmt.Println("failed to create ssh key folder: ", err)
		return false
	}

	return true
}
