package ssh

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils/paths"
	sshUtils "github.com/silentFellow/cred/internal/utils/ssh"
)

// InsertCmd represents the {cred ssh insert <filepath>} command
var InsertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert a new ssh entry",
	Long: `The insert command allows you to add a new ssh entry to the ssh store.
You will be prompted to enter and confirm the ssh, which will be stored securely.
If the entry already exists, you will be asked whether you want to overwrite it.

Examples:
ssh insert <key-name> --public-key <key-path> --private-key <key-path> --connection <connection-string>`,
	Run: func(cmd *cobra.Command, args []string) {
		usage := "ssh insert <key-name> --public-key <key-path> --private-key <key-path> --connection <connection-string>"

		sshStore := config.Constants.SshPath
		if len(args) < 1 {
			fmt.Println("invalid usage, expected: ", usage)
			return
		}

		// create a folder for the key
		path := args[0]
		keyFullPath := paths.BuildPath(sshStore, path)
		if !prepareKeyFolder(keyFullPath) {
			return
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
			removeCreated(keyFullPath)
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
				removeCreated(keyFullPath)
				return
			}

			if info.IsDir() {
				fmt.Printf("path %s is a directory, not a file\n", key.path)
				removeCreated(keyFullPath)
				return
			}

			fullPath := paths.BuildPath(keyFullPath, key.fileName)
			content := ""
			file, err := os.ReadFile(key.path)
			if err != nil {
				fmt.Printf("failed to read file %s: %v\n", key.path, err)
				removeCreated(keyFullPath)
				return
			}
			content = string(file)

			if !sshUtils.ValidateKey(content, key.keyType) {
				fmt.Printf("invalid %v ssh key\n", key.keyType)
				removeCreated(keyFullPath)
				return
			}

			if err := gpgcrypt.AddFile(fullPath, content, false); err != nil {
				fmt.Println("failed to insert ssh keys: ", err)
				removeCreated(keyFullPath)
				return
			}
		}

		if connectionString != "" {
			fullPath := paths.BuildPath(keyFullPath, "connection.gpg")
			connectionContent := connectionString

			if err := gpgcrypt.AddFile(fullPath, connectionContent, false); err != nil {
				fmt.Println("failed to insert ssh keys: ", err)
				removeCreated(keyFullPath)
				return
			}
		}

		fmt.Println("ssh inserted successfully")
	},
}

func init() {
	InsertCmd.Flags().String("public-key", "", "public key file path")
	InsertCmd.Flags().String("private-key", "", "private key file path")
	InsertCmd.Flags().String("connection-string", "", "connection string")
}
