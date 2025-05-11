package sshUtils

import "github.com/silentFellow/cred/internal/utils"

func CheckSshExists() bool {
	cmd := utils.SetCmd(utils.CmdConfig{}, "ssh", "-V")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
