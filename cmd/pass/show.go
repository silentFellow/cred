package pass

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/config"
	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils"
)

// pass/showCmd represents the pass/show command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays the stored password",
	Long: `The show command retrieves and displays the stored password for a given account.
Usage examples:

pass show <account_name>`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "pass show <filename>"

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

		fmt.Println(decryped)
	},
}
