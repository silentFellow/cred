package pass

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils"
)

// GenerateCmd represents the pass/generate command
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new password and store it securely",
	Long: `The generate command creates a new password of specified length and stores it securely in the password store.
You can specify the length of the password using the -l flag. If the file already exists, you will be prompted to overwrite it.

Examples:
  pass generate mypassword -l 16
  pass generate anotherpassword -l 24`,
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		usage := "pass generate <filename> [flags: -l {length}]"

		passStore := config.Constants.PassPath
		if len(args) < 1 {
			fmt.Printf("Invalid usage: %v\n", usage)
			return
		}
		path := args[0]
		fullPath := fmt.Sprintf("%v/%v.gpg", passStore, path)

		generatedPassword := utils.GenerateRandom(length)

		if !utils.CheckPathExists(fullPath) {
			if err := utils.AddToPath(fullPath, generatedPassword, true); err != nil {
				fmt.Println("Failed to insert password: ", err)
				return
			}
			fmt.Println("Password inserted successfully, copied to clipboard")
			return
		}

		var choice string
		fmt.Print("The file already exists. Do you want to overwrite it? (y/n): ")
		fmt.Scanln(&choice)

		if choice == "y" || choice == "Y" {
			if err := os.RemoveAll(fullPath); err != nil {
				fmt.Println("Failed to remove the file: ", err)
			}

			if err := utils.AddToPath(fullPath, generatedPassword, true); err != nil {
				fmt.Println("Failed to insert password: ", err)
				return
			}
			fmt.Println("Password inserted successfully, copied to clipboard")
		}
	},
}

func init() {
	GenerateCmd.Flags().IntP("length", "l", 12, "length of the generated password")
}
