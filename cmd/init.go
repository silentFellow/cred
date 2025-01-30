package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils/git"
	"github.com/silentFellow/cred/internal/utils/paths"
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
			fmt.Printf("invalid usage: %v\n", usage)
			return
		}

		gpgKey := args[0]
		if !gpgcrypt.CheckKeyValidity(gpgKey) {
			fmt.Printf("invalid GPG key, %v\n", usage)
			return
		}

		gpgKey, err := gpgcrypt.GetKeyFpr(gpgKey)
		if err != nil {
			fmt.Println("failed to get the key:", err)
			return
		}

		storePath := config.Constants.StorePath

		// new store
		if !paths.CheckPathExists(storePath) {
			if err := initStore(gpgKey); err != nil {
				fmt.Println("failed to initiate store:", err)
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
				fmt.Println("failed to remove store:", err)
			}

			if err := initStore(gpgKey); err != nil {
				fmt.Println("failed to initiate store:", err)
			}
		}
	},
}

func initStore(gpgid string) error {
	storeDirectoriesPaths := []string{
		config.Constants.StorePath,
		config.Constants.PassPath,
		config.Constants.EnvPath,
	}

	for _, path := range storeDirectoriesPaths {
		if err := os.MkdirAll(path, 0700); err != nil {
			return err
		}
	}

	// on success
	file, err := os.Create(paths.BuildPath(config.Constants.StorePath, ".gpg-id"))
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
	initCmd.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
		if config.Config.AutoGit {
			return git.AutoGit(cmd)
		}

		return nil
	}

	rootCmd.AddCommand(initCmd)
}
