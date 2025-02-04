package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	fscopy "github.com/silentFellow/cred/internal/utils/copy"
	"github.com/silentFellow/cred/internal/utils/git"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the credential store to a new GPG key",
	Long: `The migrate command allows you to re-encrypt your credential store with a new GPG key.
This operation will create a backup of your current store and re-encrypt all files with the new key.

	Example Usage: cred migrate <new-gpg-key-id>`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "cred migrate <new-gpg-key-id>"

		if len(args) < 1 {
			fmt.Println("invalid usage: ", usage)
			return
		}

		storePath := config.Constants.StorePath
		originalKey := config.Constants.GpgKey
		newKey := args[0]

		if !gpgcrypt.CheckKeyValidity(newKey) {
			fmt.Println("invalid key")
			return
		}

		newKey, err := gpgcrypt.GetKeyFpr(newKey)
		if err != nil {
			fmt.Println("failed to get the key:", err)
			return
		}

		var choice string
		fmt.Print("WARNING: This operation will modify the store. Do you want to continue? (y/n): ")
		fmt.Scanln(&choice)

		if strings.ToLower(choice) != "y" {
			return
		}

		tempDir, err := os.MkdirTemp("", "cred-store-migrate-")
		if err != nil {
			fmt.Println("failed to create temporary directory: ", err)
			return
		}
		defer os.RemoveAll(tempDir) // Cleanup temporary directory

		if err := fscopy.Copy(storePath, tempDir); err != nil {
			fmt.Println("failed to copy store: ", err)
			return
		}

		gpgFile, err := os.Create(filepath.Join(tempDir, ".gpg-id"))
		if err != nil {
			fmt.Printf("failed to create the .gpg-id file in temporary directory: %v\n", err)
			return
		}
		defer gpgFile.Close()

		if _, err := gpgFile.WriteString(newKey); err != nil {
			fmt.Printf("failed to write the new GPG key to the .gpg-id file: %v\n", err)
			return
		}

		// recrypt the store
		err = filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("failed to walk through the directory: %v", err)
			}

			// skip directories
			if info.IsDir() {
				return nil
			}

			relPath, err := filepath.Rel(tempDir, path)
			if err != nil {
				return fmt.Errorf("failed to calculate relative path: %v", err)
			}

			if strings.HasPrefix(relPath, "pass") || strings.HasPrefix(relPath, "env") {
				destPath := filepath.Join(tempDir, relPath)
				if err := gpgcrypt.Recrypt(destPath, originalKey, newKey); err != nil {
					fmt.Printf("failed to recrypt file %s: %v\n", relPath, err)
				}
			}

			return nil
		})
		if err != nil {
			fmt.Printf("failed to recrypt the store: %v\n", err)
			return
		}

		if err := os.RemoveAll(storePath); err != nil {
			fmt.Println("failed to remove store: ", err)
			return
		}

		if err := fscopy.Copy(tempDir, storePath); err != nil {
			fmt.Println("failed to copy store: ", err)
			return
		}

		fmt.Println("store migrated successfully")
	},
}

func init() {
	migrateCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if !gpgcrypt.CheckKeyExists() {
			cmd.SilenceUsage = true
			return fmt.Errorf("GPG key not found, try [cred init <gpg-key-id>]")
		}

		if !gpgcrypt.CheckKeyValidity(config.Constants.GpgKey) {
			cmd.SilenceUsage = true
			return fmt.Errorf("invalid GPG key, try [cred init <gpg-key-id>]")
		}

		return nil
	}

	migrateCmd.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
		if config.Config.AutoGit {
			return git.AutoGit(cmd)
		}

		return nil
	}

	rootCmd.AddCommand(migrateCmd)
}
