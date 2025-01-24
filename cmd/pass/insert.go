package pass

import (
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils"
)

// InsertCmd represents the pass/insert command
var InsertCmd = &cobra.Command{
	Use:   "insert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "pass insert <filename> [flags: -l {length}]"

		passStore := config.Constants.PassPath
		if len(args) < 1 {
			fmt.Printf("Invalid usage: %v\n", usage)
			return
		}
		path := args[0]
		fullPath := fmt.Sprintf("%v/%v.gpg", passStore, path)

		if utils.CheckPathExists(fullPath) {
			var choice string
			fmt.Print("The file already exists. Do you want to overwrite it? (y/n): ")
			fmt.Scanln(&choice)

			if choice != "y" && choice != "Y" {
				return
			}
		}

		fmt.Print("Enter password (input will be hidden): ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("Failed to read password: ", err)
		}
		fmt.Println()

		fmt.Print("Enter confirm password (input will be hidden): ")
		byteConfirmPassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("Failed to read confirm password: ", err)
		}
		fmt.Println()

		password, confirmPassword := string(bytePassword), string(byteConfirmPassword)

		if password == "" || (password != confirmPassword) {
			fmt.Println("Password don't match (or) Invalid password")
			return
		}

		if utils.CheckPathExists(fullPath) {
			if err := os.RemoveAll(fullPath); err != nil {
				fmt.Println("Failed to remove the file: ", err)
				return
			}
		}
		addToPath(fullPath, password, false)
	},
}
