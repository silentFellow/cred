package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/config"
	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils"
)

// quickSetupCmd represents the quickSetup command
var quickSetupCmd = &cobra.Command{
	Use:   "quick-setup",
	Short: "Creates a GPG ID and initiates store usage credentials",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

	Usage: cred quick-setup`,
	Run: func(cmd *cobra.Command, args []string) {
		var uname, email string

		fmt.Print("Enter your username: ")
		fmt.Scanln(&uname)

		fmt.Print("Enter your email: ")
		fmt.Scanln(&email)

		if strings.Trim(uname, " ") == "" || strings.Trim(email, " ") == "" {
			fmt.Println("Invalid input")
			return
		}

		if err := gpgcrypt.GenerateKey(uname, email); err != nil {
			fmt.Println("Failed to generate key, ", err)
			return
		}

		keyID, err := gpgcrypt.GetKeyFpr(uname)
		if err != nil {
			fmt.Println("Failed to get the key, ", err)
			return
		}

		if err := gpgcrypt.AddSubKey(keyID); err != nil {
			fmt.Println("Failed to add subkey, ", err)
			return
		}

		if err := gpgcrypt.ModifyTrust(keyID); err != nil {
			fmt.Println("Failed to modify trust, ", err)
			return
		}

		storePath := config.Constants.StorePath

		// new store
		if !utils.CheckPathExists(storePath) {
			if err := initStore(keyID); err != nil {
				fmt.Println("Failed to initiate store, ", err)
			}
			return
		}

		// overwrite existing store
		var choice string
		fmt.Printf(
			"The store already exists at %s. Do you want to overwrite it? (y/n): ",
			storePath,
		)
		fmt.Scanln(&choice)

		if strings.ToLower(choice) == "y" {
			if err := os.RemoveAll(storePath); err != nil {
				fmt.Println("Failed to remove store, ", err)
			}

			if err := initStore(keyID); err != nil {
				fmt.Println("Failed to initiate store, ", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(quickSetupCmd)
}
