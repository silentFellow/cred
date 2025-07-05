package pass

import (
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/silentFellow/cred/config"
	"github.com/silentFellow/cred/internal/core"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils/paths"
)

// InsertCmd represents the {cred pass insert <filepath>} command
var InsertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert a new password entry",
	Long: `The insert command allows you to add a new password entry to the password store.
You will be prompted to enter and confirm the password, which will be stored securely.
If the entry already exists, you will be asked whether you want to overwrite it.

Examples:
  pass insert <filepath>`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "pass insert <filename>"

		passStore := config.Constants.PassPath
		if len(args) < 1 {
			fmt.Println("invalid usage, expected: ", usage)
			return
		}
		filePath := args[0] + ".gpg"
		fullPath := paths.BuildPath(passStore, filePath)

		if paths.CheckPathExists(fullPath) {
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
			fmt.Println("failed to read password: ", err)
		}
		fmt.Println()

		fmt.Print("enter confirm password (input will be hidden): ")
		byteConfirmPassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("failed to read confirm password: ", err)
		}
		fmt.Println()

		password, confirmPassword := string(bytePassword), string(byteConfirmPassword)

		if password == "" || (password != confirmPassword) {
			fmt.Println("passwords do not match or invalid input")
			return
		}

		if paths.CheckPathExists(fullPath) {
			if err := os.RemoveAll(fullPath); err != nil {
				fmt.Println("failed to remove the file: ", err)
				return
			}
		}

		if err := gpgcrypt.AddFile(fullPath, password, true); err != nil {
			fmt.Println("failed to insert password: ", err)
			return
		}
		fmt.Println("password inserted successfully")

		isEditor, _ := cmd.Flags().GetBool("editor")
		if isEditor {
			core.EditLogic("pass", append([]string{filePath}, args[1:]...))
		}
	},
}

func init() {
	InsertCmd.Flags().
		BoolP("editor", "e", false, "open password in editor for editing extra details after insertion")
}
