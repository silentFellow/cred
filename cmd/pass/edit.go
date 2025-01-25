package pass

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/config"
	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils"
)

// EditCmd represents the pass/edit command
var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a password entry",
	Long: `Edit a password entry in the password store.

This command allows you to edit an existing password entry in your password store.
Usage:
  pass edit <filename>`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "pass edit <filename>"

		passStore := config.Constants.PassPath
		if len(args) < 1 {
			fmt.Printf("Invalid usage: %v\n", usage)
			return
		}
		path := args[0]
		fullPath := fmt.Sprintf("%v/%v", passStore, path)

		if !utils.CheckPathExists(fullPath) {
			fmt.Println("Password not found")
			return
		}

		if utils.GetPathType(fullPath) != "file" {
			fmt.Println("Invalid file format, only file is allowed")
			return
		}

		originalContent, err := gpgcrypt.Decrypt(fullPath)
		if err != nil {
			fmt.Println("Error decrypting file")
			return
		}

		tempFile, err := os.CreateTemp("", "pass-edit-*.tmp")
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

		if err := utils.AddToPath(fullPath, updatedContent, true); err != nil {
			fmt.Println("Failed to update password: ", err)
			return
		}
		fmt.Println("Password updated successfully")
	},
}
