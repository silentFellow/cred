package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred-store/config"
	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	fscopy "github.com/silentFellow/cred-store/internal/utils/copy"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "cred migrate <new-gpg-key-id>"

		if len(args) < 1 {
			fmt.Printf("Invalid usage: %v\n", usage)
			return
		}

		storePath := config.Constants.StorePath
		originalKey := config.Constants.GpgKey
		newKey := args[0]

		if !gpgcrypt.CheckKeyValidity(newKey) {
			fmt.Println("Invalid key")
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
			fmt.Println("Failed to create temporary directory: ", err)
			return
		}
		defer os.RemoveAll(tempDir) // Cleanup temporary directory

		if err := fscopy.Copy(storePath, tempDir); err != nil {
			fmt.Println("Failed to copy store: ", err)
			return
		}

		gpgFile, err := os.Create(filepath.Join(tempDir, ".gpg-id"))
		if err != nil {
			fmt.Printf("Failed to create the .gpg-id file in temporary directory: %v\n", err)
			return
		}
		defer gpgFile.Close()

		if _, err := gpgFile.WriteString(newKey); err != nil {
			fmt.Printf("Failed to write the new GPG key to the .gpg-id file: %v\n", err)
			return
		}

		// recrypt the store
		err = filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("Failed to walk through the directory: %v", err)
			}

			// skip directories
			if info.IsDir() {
				return nil
			}

			relPath, err := filepath.Rel(tempDir, path)
			if err != nil {
				return fmt.Errorf("Failed to calculate relative path: %v", err)
			}

			if relPath != ".gpg-id" {
				destPath := filepath.Join(tempDir, relPath)
				if err := gpgcrypt.Recrypt(destPath, originalKey, newKey); err != nil {
					return fmt.Errorf("Failed to recrypt data: %v", err)
				}
			}

			return nil
		})
		if err != nil {
			fmt.Printf("Failed to recrypt the store: %v\n", err)
			return
		}

		// Backup original store
		backupPath := storePath + ".backup"
		if err := fscopy.Copy(storePath, backupPath); err != nil {
			fmt.Printf("Failed to backup original store: %v\n", err)
			return
		}

		if err := os.RemoveAll(storePath); err != nil {
			fmt.Println("Failed to remove store: ", err)
			return
		}

		if err := fscopy.Copy(tempDir, storePath); err != nil {
			fmt.Println("Failed to copy store: ", err)
			return
		}

		fmt.Println("Store migrated successfully")
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
			return fmt.Errorf("Invalid GPG key, try [cred init <gpg-key-id>]")
		}

		return nil
	}

	rootCmd.AddCommand(migrateCmd)
}
