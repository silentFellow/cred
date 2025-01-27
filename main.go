package main

import (
	"fmt"

	"github.com/silentFellow/cred-store/cmd"
	"github.com/silentFellow/cred-store/internal/utils"
)

func main() {
	gpgCmd := utils.SetCmd("", utils.CmdIOConfig{}, "gpg", "--version")
	if _, err := gpgCmd.CombinedOutput(); err != nil {
		fmt.Println("GPG is not installed. Please install GPG to use this tool")
		return
	}

	cmd.Execute()
}
