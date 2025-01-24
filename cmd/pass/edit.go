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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "pass insert <filename>"

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

		addToPath(fullPath, updatedContent, false)
	},
}
