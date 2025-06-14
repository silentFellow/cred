package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils/paths"
)

// SetCmd represents the {cred env set <filepath>} command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets environment variables from the cred-store into a .env file",
	Long: `Retrieve environment variables for a specified file from the cred-store and write them to a local .env file.

You provide the file name as an argument. The command fetches the environment variables associated with that file from the cred-store and writes them into a new or existing .env file in your current directory.

	Example usage: cred env set <filepath>`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "cred env set <filename>"

		basePath := config.Constants.EnvPath

		if len(args) < 1 || len(args) > 1 {
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
			fmt.Println("invalid file format, only file is allowed")
			return
		}

		destFile := ".env"
		if paths.CheckPathExists(destFile) {
			var choice string
			fmt.Print(
				".env file already exists in the current folder. Do you want to overwrite it? (y/n): ",
			)
			fmt.Scanln(&choice)

			if strings.ToLower(choice) != "y" {
				return
			}
		}

		decrypted, err := gpgcrypt.Decrypt(fullPath)
		if err != nil {
			fmt.Printf("decrypting file %v failed: %v\n", path, err)
			return
		}

		if err := os.WriteFile(destFile, []byte(decrypted), 0644); err != nil {
			fmt.Println("failed to write values to .env: ", err)
			return
		}

		fmt.Println("successfully set environment variables to .env file")
		fmt.Println("make sure .env is listed in your .gitignore to avoid committing secrets.")
	},
}
