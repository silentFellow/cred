package core

import (
	"fmt"
	"os"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils"
	"github.com/silentFellow/cred/internal/utils/paths"
)

func EditLogic(
	cmdType string,
	args []string,
) {
	usage := fmt.Sprintf("cred %v edit <filename>", cmdType)

	var basePath string
	switch cmdType {
	case "pass":
		basePath = config.Constants.PassPath
	case "env":
		basePath = config.Constants.EnvPath
	}

	if len(args) < 1 {
		fmt.Println("invalid usage, expected: ", usage)
		return
	}

	path := args[0]
	fullPath := paths.BuildPath(basePath, path)

	if !paths.CheckPathExists(fullPath) {
		fmt.Printf("%v not found\n", path)
		return
	}

	if paths.GetPathType(fullPath) != "file" {
		fmt.Println("invalid file format, only files are allowed")
		return
	}

	originalContent, err := gpgcrypt.Decrypt(fullPath)
	if err != nil {
		if err.Error() == "wrong passphrase" {
			fmt.Println("Invalid passphrase, please try again")
		} else {
			fmt.Println("invalid file format, only files are allowed")
		}
		return
	}

	tempFile, err := os.CreateTemp("", fmt.Sprintf("%v-edit-*.tmp", cmdType))
	if err != nil {
		fmt.Println("creating temp file failed: ", err)
		return
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.WriteString(originalContent); err != nil {
		fmt.Println("writing to temp file failed: ", err)
		return
	}

	editorCmd := utils.SetCmd(
		utils.CmdConfig{IsStdin: true, IsStdout: true, IsStderr: true},
		config.Config.Editor,
		tempFile.Name(),
	)

	if err := editorCmd.Run(); err != nil {
		fmt.Println("opening editor failed: ", err)
		return
	}

	updatedContentBytes, err := os.ReadFile(tempFile.Name())
	if err != nil {
		fmt.Println("reading updated contents failed: ", err)
		return
	}

	updatedContent := string(updatedContentBytes)

	if updatedContent == originalContent {
		fmt.Println("no changes detected")
		return
	}

	if err := gpgcrypt.AddFile(fullPath, updatedContent, true); err != nil {
		fmt.Println("updating file failed: ", err)
		return
	}

	fmt.Printf("%v updated successfully\n", path)
}
