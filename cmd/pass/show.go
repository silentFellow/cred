package pass

import (
	"fmt"

	"github.com/silentFellow/cred-store/config"
	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils"
	"github.com/spf13/cobra"
)

// pass/showCmd represents the pass/show command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
