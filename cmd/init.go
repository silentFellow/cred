package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

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
		if !checkKeyValidity(gpgKey) {
			fmt.Printf("Invalid GPG key, %v\n", usage)
			return
		}

		storePath := config.Constants.StorePath

		// new store
		if !utils.CheckPathExists(storePath) {
			if err := makeStore(); err != nil {
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

			if err := makeStore(); err != nil {
				fmt.Println("Failed to initiate store, ", err)
			}
		}
	},
}

func checkKeyValidity(keyId string) bool {
	cmd := exec.Command("gpg", "--list-keys", keyId)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	if strings.Contains(string(output), keyId) {
		return true
	}

	return false
}

func makeStore() error {
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

	fmt.Println("Store initiated successfully.")
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}
