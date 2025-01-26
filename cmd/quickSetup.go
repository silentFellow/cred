package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/config"
	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils/git"
	"github.com/silentFellow/cred-store/internal/utils/paths"
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
			fmt.Println("Failed to get the key: ", err)
			return
		}

		if err := gpgcrypt.AddSubKey(keyID); err != nil {
			fmt.Println("Failed to add subkey: ", err)
			return
		}

		if err := gpgcrypt.ModifyTrust(keyID); err != nil {
			fmt.Println("Failed to modify trust: ", err)
			return
		}

		if err := gpgcrypt.ExportKeys(keyID); err != nil {
			fmt.Println("Failed to export keys: ", err)
			return
		}

		storePath := config.Constants.StorePath

		// new store
		if !paths.CheckPathExists(storePath) {
			if err := initStore(keyID); err != nil {
				fmt.Println("Failed to initiate store, ", err)
			}
			return
		}

		// overwrite existing store
		fmt.Printf("The store already exists at %s.\n", storePath)
		fmt.Println("Choose an option:")
		fmt.Println("1. Migrate the store")
		fmt.Println("2. Overwrite the store")
		fmt.Println("n. Do nothing and exit")

		var choice string
		fmt.Print("Enter your choice (1/2/n): ")
		fmt.Scanln(&choice)

		switch strings.ToLower(choice) {
		case "1":
			// Migrate the store
			migrateCmd.Run(cmd, []string{keyID})
		case "2":
			// Overwrite the store
			if err := os.RemoveAll(storePath); err != nil {
				fmt.Println("Failed to remove store, ", err)
			}

			if err := initStore(keyID); err != nil {
				fmt.Println("Failed to initiate store, ", err)
			}
			fmt.Println("Store overwritten successfully.")
		case "n":
			// Exit without doing anything
			fmt.Println("No changes made. Exiting.")
		default:
			// Invalid input
			fmt.Println("Invalid choice. No changes made.")
		}
	},
}

func init() {
	quickSetupCmd.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
		if config.Constants.AutoGit {
			return git.AutoGit(cmd)
		}

		return nil
	}

	rootCmd.AddCommand(quickSetupCmd)
}
