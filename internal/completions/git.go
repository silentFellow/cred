package completions

import (
	"github.com/spf13/cobra"

	"github.com/silentFellow/cred/config"
	"github.com/silentFellow/cred/internal/utils/git"
)

// func GetStageableCompletions =

// GitCommandMap provides a map of common Git commands for GitHub
var GitCommandMap = map[string]string{
	"clone":        "Clone a repository into a new directory",
	"init":         "Create an empty Git repository or reinitialize an existing one",
	"add":          "Add file contents to the index",
	"commit":       "Record changes to the repository",
	"status":       "Show the working tree status",
	"log":          "Show commit logs",
	"pull":         "Fetch from and integrate with another repository or a local branch",
	"push":         "Update remote refs along with associated objects",
	"branch":       "List, create, or delete branches",
	"checkout":     "Switch branches or restore working tree files",
	"merge":        "Join two or more development histories together",
	"rebase":       "Reapply commits on top of another base tip",
	"remote":       "Manage set of tracked repositories",
	"fetch":        "Download objects and refs from another repository",
	"tag":          "List, create, or delete tags",
	"diff":         "Show changes between commits, commit and working tree, etc.",
	"reset":        "Reset current HEAD to the specified state",
	"rm":           "Remove files from the working tree and from the index",
	"stash":        "Stash the changes in a dirty working directory away",
	"apply":        "Apply a patch to files and/or to the index",
	"cherry-pick":  "Apply the changes introduced by some existing commits",
	"blame":        "Show what revision and author last modified each line of a file",
	"show":         "Show various types of objects",
	"describe":     "Give an object a human-readable name based on an available ref",
	"config":       "Get and set Git configuration values",
	"revert":       "Revert a commit by creating a new commit that undoes changes",
	"bisect":       "Use binary search to find the commit that introduced a bug",
	"reflog":       "Show the history of the references (HEAD, branch, etc.)",
	"gc":           "Optimize the repository by cleaning up unnecessary files",
	"clean":        "Remove untracked files from the working directory",
	"submodule":    "Manage Git submodules",
	"archive":      "Create an archive of files from a Git repository",
	"ls-tree":      "List the contents of a tree object (commit, branch, etc.)",
	"update-index": "Manually manipulate the index (e.g., add/remove files)",
}

func GetGitFileCompletion(
	cmd *cobra.Command,
	args []string,
	toComplete string,
) ([]string, cobra.ShellCompDirective) {
	stageables, err := git.GetStageable(config.Constants.StorePath)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}
	return stageables, cobra.ShellCompDirectiveDefault
}
