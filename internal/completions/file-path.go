package completions

import (
	"os"
	"path/filepath"
)

// GetFilePathSuggestions retrieves all file paths and directories for autocompletion
func GetFilePathSuggestions(basePath string) []string {
	var suggestions []string

	// Walk through all files and directories under basePath
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		// Convert full path to relative path
		relativePath, _ := filepath.Rel(basePath, path)

		// Skip the base path itself
		if relativePath == "." {
			return nil
		}

		// Append directories with a trailing slash
		if info.IsDir() {
			suggestions = append(suggestions, relativePath)
		} else {
			suggestions = append(suggestions, relativePath)
		}

		return nil
	})
	// If an error occurs during traversal, return an empty slice
	if err != nil {
		return nil
	}

	return suggestions
}
