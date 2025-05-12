package ssh

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	fscopy "github.com/silentFellow/cred/internal/utils/copy"
	"github.com/silentFellow/cred/internal/utils/paths"
	sshUtils "github.com/silentFellow/cred/internal/utils/ssh"
)

// EditCmd represents the {cred ssh edit <filepath>} command
var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit a new ssh entry",
	Long: `The edit command allows you to add a new ssh entry to the ssh store.
You will be prompted to enter and confirm the ssh, which will be stored securely.
If the entry already exists, you will be asked whether you want to overwrite it.

Examples:
ssh edit <key-name> --public-key <key-path> --private-key <key-path> --connection <connection-string>`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "ssh edit <key-name> --public-key <key-path> --private-key <key-path> --connection <connection-string>`"

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

		// create a backup dir
		backupPath := keyFullPath + "_backup"
		err = fscopy.CopyDirectory(keyFullPath, backupPath)
		if err != nil {
			fmt.Println("could not create backup of SSH key: ", err)
			return
		}

		// backup function
		restoreOnError := func() {
			fmt.Println("An error occurred. Restoring from backup...")
			os.RemoveAll(keyFullPath)
			os.Rename(backupPath, keyFullPath)
		}

		// get flags
		publicKeyPath, err := cmd.Flags().GetString("public-key")
		if err != nil {
			fmt.Println("failed to get public-key flag: ", err)
			return
		}

		privateKeyPath, err := cmd.Flags().GetString("private-key")
		if err != nil {
			fmt.Println("failed to get private-key flag: ", err)
			return
		}

		connectionString, err := cmd.Flags().GetString("connection-string")
		if err != nil {
			fmt.Println("failed to get connection-string flag: ", err)
			return
		}

		if publicKeyPath == "" && privateKeyPath == "" && connectionString == "" {
			fmt.Println("no files provided")
			return
		}

		// create files based on it
		type keyFormat struct {
			keyType  string
			fileName string
			path     string
		}
		keys := []keyFormat{
			{
				keyType:  "public",
				fileName: "public.gpg",
				path:     publicKeyPath,
			},
			{
				keyType:  "private",
				fileName: "private.gpg",
				path:     privateKeyPath,
			},
		}

		for _, key := range keys {
			if key.path == "" {
				continue
			}

			info, err := os.Stat(key.path)
			if err != nil {
				if os.IsNotExist(err) {
					fmt.Printf("file %s does not exist\n", key.path)
				} else {
					fmt.Printf("failed to access file %s: %v\n", key.path, err)
				}
				restoreOnError()
				return
			}

			if info.IsDir() {
				fmt.Printf("path %s is a directory, not a file\n", key.path)
				restoreOnError()
				return
			}

			fullPath := paths.BuildPath(keyFullPath, key.fileName)
			content := ""
			file, err := os.ReadFile(key.path)
			if err != nil {
				fmt.Printf("failed to read file %s: %v\n", key.path, err)
				restoreOnError()
				return
			}
			content = string(file)

			if !sshUtils.ValidateKey(content, key.keyType) {
				fmt.Printf("invalid %v ssh key\n", key.keyType)
				restoreOnError()
				return
			}

			if err := gpgcrypt.AddFile(fullPath, content, false); err != nil {
				fmt.Println("failed to insert ssh keys: ", err)
				restoreOnError()
				return
			}
		}

		if connectionString != "" {
			fullPath := paths.BuildPath(keyFullPath, "connection.gpg")
			connectionContent := connectionString

			if err := gpgcrypt.AddFile(fullPath, connectionContent, false); err != nil {
				fmt.Println("failed to edit ssh keys: ", err)
				restoreOnError()
				return
			}
		}

		if err := os.RemoveAll(backupPath); err != nil {
			fmt.Println("Failed to remove the backup")
		}
		fmt.Println("SSH edited successfully")
	},
}

func init() {
	EditCmd.Flags().String("public-key", "", "public key file path")
	EditCmd.Flags().String("private-key", "", "private key file path")
	EditCmd.Flags().String("connection-string", "", "connection string")
}
