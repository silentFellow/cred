package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/config"
	"github.com/silentFellow/cred-store/internal/utils"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init <gpg-key>",
	Short: "Initialize the credential store",
	Long: `Initialize the credential store with the necessary settings.
This command sets up the environment required for storing
and managing credentials securely. Note that you must provide a GPG key as an argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "cred init <gpg-key-id>"

		if len(args) < 1 {
			fmt.Printf("Invalid usage: %v\n", usage)
			return
		}

		gpgKey := args[0]
		if !utils.CheckKeyValidity(gpgKey) {
			fmt.Printf("Invalid GPG key, %v\n", usage)
			return
		}

		storePath := config.Constants.StorePath

		// new store
		if !utils.CheckPathExists(storePath) {
			if err := initStore(gpgKey); err != nil {
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

		if choice == "y" || choice == "Y" {
			if err := os.RemoveAll(storePath); err != nil {
				fmt.Println("Failed to remove store, ", err)
			}

			if err := initStore(gpgKey); err != nil {
				fmt.Println("Failed to initiate store, ", err)
			}
		}
	},
}

func initStore(gpgid string) error {
	paths := []string{
		config.Constants.StorePath,
		config.Constants.PassPath,
		config.Constants.EnvPath,
	}

	for _, path := range paths {
		if err := os.MkdirAll(path, 0700); err != nil {
			return err
		}
	}

	// on success
	file, err := os.Create(fmt.Sprintf("%v/.gpg-id", config.Constants.StorePath))
	defer file.Close()

	if err != nil {
		return err
	}

	if _, err := file.WriteString(gpgid); err != nil {
		return err
	}

	fmt.Println("Store initiated successfully.")
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}
