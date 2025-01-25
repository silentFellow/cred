package core

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/silentFellow/cred-store/config"
	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils/paths"
)

func EditLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v edit <filename>", cmdType)

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

	if !paths.CheckPathExists(fullPath) {
		fmt.Println("Path not found")
		return
	}

	if paths.GetPathType(fullPath) != "file" {
		fmt.Println("Invalid file format, only file is allowed")
		return
	}

	originalContent, err := gpgcrypt.Decrypt(fullPath)
	if err != nil {
		fmt.Println("Error decrypting file")
		return
	}

	tempFile, err := os.CreateTemp("", fmt.Sprintf("%v-edit-*.tmp", cmdType))
	if err != nil {
		fmt.Println("Error creating temp file")
		return
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(originalContent); err != nil {
		fmt.Println("Error writing to temp file")
		return
	}

	editorCmd := exec.Command(config.Constants.Editor, tempFile.Name())
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr

	if err := editorCmd.Run(); err != nil {
		fmt.Println("Error opening editor, ", err)
		return
	}

	updatedContentBytes, err := os.ReadFile(tempFile.Name())
	if err != nil {
		fmt.Println("Failed to read the updated contents")
		return
	}

	updatedContent := string(updatedContentBytes)

	if updatedContent == originalContent {
		fmt.Println("No changes detected")
		return
	}

	if err := paths.AddToPath(fullPath, updatedContent, true); err != nil {
		fmt.Println("Failed to update file: ", err)
		return
	}
	fmt.Println("File updated successfully")
}
