package env

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils"
)

// InsertCmd represents the {cred env insert <filepath>} command
var InsertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert a new env entry",
	Long: `The insert command allows you to add a new env entry to the env store.
You will be prompted to enter and confirm the env, which will be stored securely.
If the entry already exists, you will be asked whether you want to overwrite it.

Examples:
  cred env insert <filepath>`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "cred env insert <filename>"

		basePath := config.Constants.EnvPath

		if len(args) < 1 {
			fmt.Printf("Invalid usage: %v\n", usage)
			return
		}

		path := args[0]
		fullPath := fmt.Sprintf("%v/%v", basePath, path)

		if utils.CheckPathExists(fullPath) {
			var choice string
			fmt.Print("The file already exists. Do you want to overwrite it? (y/n): ")
			fmt.Scanln(&choice)

			if strings.ToLower(choice) != "y" {
				return
			}
		}

		tempFile, err := os.CreateTemp("", "env-insert-*.tmp")
		if err != nil {
			fmt.Println("Error creating temp file")
			return
		}
		defer os.Remove(tempFile.Name())

		editorCmd := exec.Command(config.Constants.Editor, tempFile.Name())
		editorCmd.Stdin = os.Stdin
		editorCmd.Stdout = os.Stdout
		editorCmd.Stderr = os.Stderr

		if err := editorCmd.Run(); err != nil {
			fmt.Println("Error opening editor, ", err)
			return
		}

		contentBytes, err := os.ReadFile(tempFile.Name())
		if err != nil {
			fmt.Println("Failed to read the contents")
			return
		}

		content := string(contentBytes)

		if content == "" {
			fmt.Println("Invalid content")
			return
		}

		if err := utils.AddToPath(fullPath, content, true); err != nil {
			fmt.Println("Failed to update env: ", err)
			return
		}
		fmt.Println("Env updated successfully")
	},
}
