package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils"
	"github.com/silentFellow/cred/internal/utils/paths"
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
			fmt.Println("invalid usage, expected: ", usage)
			return
		}

		path := args[0] + ".gpg"
		fullPath := paths.BuildPath(basePath, path)

		if paths.CheckPathExists(fullPath) {
			var choice string
			fmt.Print("The file already exists. Do you want to overwrite it? (y/n): ")
			fmt.Scanln(&choice)

			if strings.ToLower(choice) != "y" {
				return
			}
		}

		tempFile, err := os.CreateTemp("", "env-insert-*.tmp")
		if err != nil {
			fmt.Println("error creating temporary file: ", err)
			return
		}
		defer os.Remove(tempFile.Name())

		editorCmd := utils.SetCmd(
			"",
			utils.CmdIOConfig{IsStdin: true, IsStdout: true, IsStderr: true},
			config.Config.Editor,
			tempFile.Name(),
		)

		if err := editorCmd.Run(); err != nil {
			fmt.Println("opening editor failed: ", err)
			return
		}

		contentBytes, err := os.ReadFile(tempFile.Name())
		if err != nil {
			fmt.Println("Failed to read the contents: ", err)
			return
		}

		content := string(contentBytes)

		if content == "" {
			fmt.Println("invalid content")
			return
		}

		if err := gpgcrypt.AddFile(fullPath, content, true); err != nil {
			fmt.Println("failed to update env: ", err)
			return
		}

		fmt.Printf("%v inserted successfully\n", path)
	},
}
