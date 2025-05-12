package ssh

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
	gpgcrypt "github.com/silentFellow/cred/internal/gpg-crypt"
	"github.com/silentFellow/cred/internal/utils"
	"github.com/silentFellow/cred/internal/utils/paths"
	sshUtils "github.com/silentFellow/cred/internal/utils/ssh"
)

// GenerateCmd represents the {cred ssh generate} command
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a new SSH key pair and save it securely",
	Long: `The generate command allows you to create a new SSH key pair and store it
in the credential store. You can optionally add a connection string during generation.

Example:
  cred ssh generate <key-name> --connection <connection-string>`,

	Run: func(cmd *cobra.Command, args []string) {
		usage := "ssh generate <key-name> --connection <connection-string>"

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

		// create a folder to store the raw keys
		sshRawDirLocation := paths.BuildPath(config.Constants.Home, path)
		if err := os.MkdirAll(sshRawDirLocation, 0700); err != nil {
			fmt.Println("failed to create raw ssh key folder: ", err)
			return
		}

		// generate raw ssh keys
		sshRawKeyLocation := paths.BuildPath(sshRawDirLocation, "ssh_raw")
		sshKeyGenCmd := utils.SetCmd(
			utils.CmdConfig{IsStdin: true, IsStdout: true, IsStderr: true},
			"ssh-keygen",
			"-t",
			"rsa",
			"-b",
			"4096",
			"-f",
			sshRawKeyLocation,
		)

		// store public key in a file
		if err := sshKeyGenCmd.Run(); err != nil {
			fmt.Println("failed to generate ssh keys: ", err)
			removeCreated(keyFullPath)
			return
		}

		// remove raw keys
		defer func() {
			if err := os.RemoveAll(sshRawDirLocation); err != nil {
				fmt.Printf(
					"failed to remove raw keys(remove it manually at %v): %v\n",
					sshRawDirLocation,
					err,
				)
			}
		}()

		// Encrypt and store the public key
		publicKeyConent, err := os.ReadFile(sshRawKeyLocation + ".pub")
		if err != nil {
			fmt.Println("failed to read public key file: ", err)
			removeCreated(keyFullPath)
			return
		}
		publicKeyPath := paths.BuildPath(keyFullPath, "public.gpg")

		if err := gpgcrypt.AddFile(publicKeyPath, string(publicKeyConent), false); err != nil {
			fmt.Println("failed to insert ssh keys: ", err)
			removeCreated(keyFullPath)
			return
		}

		// store private key in a file
		privateKeyConent, err := os.ReadFile(sshRawKeyLocation)
		if err != nil {
			fmt.Println("failed to read private key file: ", err)
			removeCreated(keyFullPath)
			return
		}
		privateKeyPath := paths.BuildPath(keyFullPath, "private.gpg")

		if err := gpgcrypt.AddFile(privateKeyPath, string(privateKeyConent), false); err != nil {
			fmt.Println("failed to generate ssh keys: ", err)
			removeCreated(keyFullPath)
			return
		}

		connectionString, err := cmd.Flags().GetString("connection-string")
		if err != nil {
			fmt.Println("failed to get connection-string flag: ", err)
			return
		}

		if connectionString != "" {
			fullPath := paths.BuildPath(keyFullPath, "connection.gpg")
			connectionContent := connectionString

			if err := gpgcrypt.AddFile(fullPath, connectionContent, false); err != nil {
				fmt.Println("failed to generate ssh keys: ", err)
				removeCreated(keyFullPath)
				return
			}
		}

		fmt.Println("ssh generated successfully")
	},
}

func init() {
	GenerateCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if !sshUtils.CheckSshExists() {
			cmd.SilenceUsage = true
			return fmt.Errorf("SSH is not installed")
		}

		return nil
	}

	GenerateCmd.Flags().String("connection-string", "", "connection string")
}
