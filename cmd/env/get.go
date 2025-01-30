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

// GetCmd represents the {cred env get} command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve and store environment variables",
	Long: `Retrieve environment variables from a specified file or default .env files and store them securely.

You can specify the file containing the environment variables using the -f flag. If the -f flag is not provided, the command will look for the following files in order: .env, .env.local, .env.development, .env.production, .env.test.

	Example usage: cred env get [flags: -f (filepath)]`,
	Run: func(cmd *cobra.Command, args []string) {
		flagFile, _ := cmd.Flags().GetString("file")
		checkFiles := []string{
			".env",
			".env.local",
			".env.development",
			".env.production",
			".env.test",
			flagFile,
		}

		foundFile := getEnvFile(flagFile, checkFiles)

		if foundFile == "" {
			fmt.Printf("no .env file found (checked: %v)\n", checkFiles)
			return
		}

		file, err := os.ReadFile(foundFile)
		if err != nil {
			fmt.Println("failed to read file: ", foundFile)
			return
		}

		fileContent := string(file)

		var filename string
		fmt.Print("enter file name: ")
		fmt.Scanln(&filename)

		if strings.Trim(filename, " ") == "" {
			fmt.Println("invalid path")
			return
		}

		fullPath := paths.BuildPath(config.Constants.EnvPath, filename+".gpg")

		exists := paths.CheckPathExists(fullPath) && paths.GetPathType(fullPath) == "file"

		if exists {
			var choice string
			fmt.Print("The file already exists. Do you want to overwrite it? (y/n): ")
			fmt.Scanln(&choice)

			if strings.ToLower(choice) != "y" {
				return
			}

			if err := os.RemoveAll(fullPath); err != nil {
				fmt.Println("failed to remove the existing file: ", err)
			}
		}

		if err := gpgcrypt.AddFile(fullPath, fileContent, false); err != nil {
			fmt.Println("failed to insert environment variables: ", err)
			return
		}

		fmt.Println("environment variables stored successfully")
	},
}

func getEnvFile(flagFile string, checkFiles []string) string {
	if flagFile != "" {
		return flagFile
	}

	foundFile := ""
	for _, file := range checkFiles {
		if _, err := os.Stat(file); err == nil {
			foundFile = file
			break
		}
	}

	return foundFile
}

func init() {
	GetCmd.Flags().StringP("file", "f", "", "path to the file containing env")
}
