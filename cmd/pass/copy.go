package pass

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/config"
	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils"
)

// CopyCmd represents the pass/show command
var CopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copies the stored password to system clipboard",
	Long: `The show command retrieves and copies the stored password for a given account.
Usage examples:

pass show <account_name>`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "pass copy <filename>"

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

		decryped, err := gpgcrypt.Decrypt(fullPath)
		if err != nil {
			fmt.Println("Error decrypting file")
			return
		}

		if err := clipboard.WriteAll(decryped); err != nil {
			fmt.Println("Error copying to clipboard")
			return
		}

		fmt.Println("Successfully copies to clipboard")
	},
}
