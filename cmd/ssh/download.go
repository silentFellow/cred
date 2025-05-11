package ssh

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils/paths"
)

// DownloadCmd represents the {cred ssh download <filepath>} command
var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download SSH key files for the specified entry",
	Long: `The download command retrieves SSH key files (public, private, and connection info)
for a given SSH entry and saves them into a 'downloads' directory.

Example:
  cred ssh download my-key-name
This will create:
  downloads/my-key-name/public.key
  downloads/my-key-name/private.key
  downloads/my-key-name/connection.txt`,

	Run: func(cmd *cobra.Command, args []string) {
		usage := "ssh download <key-name>"

		sshStore := config.Constants.SshPath
		if len(args) < 1 {
			fmt.Println("invalid usage, expected: ", usage)
			return
		}

		// create a folder for the key
		path := args[0]
		keyFullPath := paths.BuildPath(sshStore, path)

		if !paths.CheckPathExists(keyFullPath) {
			fmt.Println("SSH key entry does not exist")
			return
		}

		keyInfo, err := os.Stat(keyFullPath)
		if err != nil {
			fmt.Println("failed to retrieve SSH key info: ", err)
			return
		}

		if !keyInfo.IsDir() {
			fmt.Println("specified SSH key path must be a directory")
			return
		}

		downloadPath := paths.BuildPath(config.Constants.Download, path+"_ssh_key")
		if err := os.MkdirAll(downloadPath, 0700); err != nil {
			fmt.Println("failed to create download directory:", err)
			return
		}

		filesToCopy := map[string]string{
			"public.gpg":     path + ".pub",
			"private.gpg":    path,
			"connection.gpg": "connection",
		}

		for src, dest := range filesToCopy {
			originalFilePath := paths.BuildPath(keyFullPath, src)
			info, err := os.Stat(originalFilePath)
			if err != nil || info.IsDir() {
				continue
			}

			content, err := gpgcrypt.Decrypt(originalFilePath)
			if err != nil {
				fmt.Printf("failed to decrypt %s: %v\n", src, err)
				return
			}

			destPath := paths.BuildPath(downloadPath, dest)
			destFile, err := os.Create(destPath)
			if err != nil {
				fmt.Printf("failed to create %s: %v\n", destPath, err)
				return
			}
			defer destFile.Close()

			if _, err := destFile.WriteString(content); err != nil {
				fmt.Printf("failed to write to %s: %v\n", destPath, err)
				return
			}
		}

		fmt.Printf("SSH key download complete. Files saved in: %s\n", downloadPath)
	},
}
