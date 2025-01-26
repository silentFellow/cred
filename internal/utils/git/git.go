package git

import (
	"os"
	"os/exec"
)

func CheckGitExists() bool {
	cmd := exec.Command("git", "--version")
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

	cmd := exec.Command("git", "status")
	cmd.Dir = filepath
	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}

func HaveRemote(filepath string) bool {
	cmd := exec.Command("git", "remote", "-v")
	cmd.Dir = filepath

	output, err := cmd.CombinedOutput()
	if string(output) == "" || err != nil {
		return false
	}
	return true
}

func HaveDiff(filePath string) bool {
	cmdDiff := exec.Command("git", "diff", "--quiet")
	cmdDiff.Dir = filePath

	// Check if there are any modifications, deletions, or renames in tracked files
	output, errDiff := cmdDiff.CombinedOutput()
	if string(output) != "" || errDiff != nil {
		return true
	}

	// Check for untracked files
	cmdUntracked := exec.Command("git", "ls-files", "--others", "--exclude-standard")
	cmdUntracked.Dir = filePath

	output, errUntracked := cmdUntracked.CombinedOutput()
	if string(output) != "" || errUntracked != nil {
		return true
	}

	// Check for deleted files (files that have been deleted but not staged)
	cmdDeleted := exec.Command("git", "ls-files", "--deleted")
	cmdDeleted.Dir = filePath

	output, errDeleted := cmdDeleted.CombinedOutput()
	if string(output) != "" || errDeleted != nil {
		return true
	}

	return false
}

func InitRepo(filePath string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = filePath
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func AddFiles(filePath string) error {
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = filePath
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func CommitFiles(filePath, message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Dir = filePath
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func PushRepo(filePath string) error {
	cmd := exec.Command("git", "push")
	cmd.Dir = filePath
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
