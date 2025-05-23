package git

import (
	"os"
	"strings"

	"github.com/silentFellow/cred/internal/utils"
)

func CheckGitExists() bool {
	cmd := utils.SetCmd(utils.CmdConfig{}, "git", "--version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func IsValidGitPath(filepath string) bool {
	path, err := os.Stat(filepath)
	if err != nil || !path.IsDir() {
		return false
	}

	cmd := utils.SetCmd(utils.CmdConfig{Dir: filepath}, "git", "status")
	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}

func HaveRemote(filepath string) bool {
	cmd := utils.SetCmd(
		utils.CmdConfig{Dir: filepath},
		"git",
		"remote",
		"-v",
	)

	output, err := cmd.CombinedOutput()
	if string(output) == "" || err != nil {
		return false
	}
	return true
}

func HaveDiff(filePath string) bool {
	cmdDiff := utils.SetCmd(
		utils.CmdConfig{Dir: filePath},
		"git",
		"diff",
		"--quiet",
	)

	// Check if there are any modifications, deletions, or renames in tracked files
	output, errDiff := cmdDiff.CombinedOutput()
	if string(output) != "" || errDiff != nil {
		return true
	}

	// Check for untracked files
	cmdUntracked := utils.SetCmd(
		utils.CmdConfig{Dir: filePath},
		"git",
		"ls-files",
		"--others",
		"--exclude-standard",
	)

	output, errUntracked := cmdUntracked.CombinedOutput()
	if string(output) != "" || errUntracked != nil {
		return true
	}

	// Check for deleted files (files that have been deleted but not staged)
	cmdDeleted := utils.SetCmd(utils.CmdConfig{Dir: filePath}, "git", "ls-files", "--deleted")

	output, errDeleted := cmdDeleted.CombinedOutput()
	if string(output) != "" || errDeleted != nil {
		return true
	}

	return false
}

func InitRepo(filePath string) error {
	cmd := utils.SetCmd(
		utils.CmdConfig{IsStdout: true, IsStderr: true, Dir: filePath},
		"git",
		"init",
	)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func AddFiles(filePath string) error {
	cmd := utils.SetCmd(utils.CmdConfig{IsStderr: true, Dir: filePath}, "git", "add", ".")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func CommitFiles(filePath, message string) error {
	cmd := utils.SetCmd(
		utils.CmdConfig{IsStdout: true, IsStderr: true, Dir: filePath},
		"git",
		"commit",
		"-m",
		message,
	)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func PushRepo(filePath string) error {
	cmd := utils.SetCmd(
		utils.CmdConfig{IsStdout: true, IsStderr: true, Dir: filePath},
		"git",
		"push",
	)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func GetStageable(filePath string) ([]string, error) {
	// trackable files
	cmd := utils.SetCmd(utils.CmdConfig{Dir: filePath}, "git", "status", "--porcelain")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")

	files := []string{}

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) >= 2 {
			status := strings.ToLower(fields[0])
			if status != "d" && status != "am" && status != "a" {
				files = append(files, fields[1])
			}
		}
	}

	return files, nil
}
