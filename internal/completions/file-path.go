package completions

import (
	"os"
	"path/filepath"
	"sort"
)

// FilePathSuggestionOptions defines the configuration for generating file path suggestions.
type FilePathSuggestionOptions struct {
	BasePath   string // The base directory to start scanning from
	AllowDirs  bool   // Whether to include directories in the suggestions
	AllowFiles bool   // Whether to include files in the suggestions
}

// GetFilePathSuggestions returns a sorted list of relative file or directory paths
// from the given base path, based on the configuration.
// If both AllowDirs and AllowFiles are false, the result will be empty.
func GetFilePathSuggestions(config FilePathSuggestionOptions) []string {
	var suggestions []string

	// Walk through all files and directories under basePath
	err := filepath.Walk(config.BasePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		// Skip the base path itself
		if config.BasePath == path {
			return nil
		}

		// Apply filtering rules based on config
		if (!config.AllowDirs && info.IsDir()) || (!config.AllowFiles && !info.IsDir()) {
			return nil
		}

		// Convert full path to relative path
		relativePath, err := filepath.Rel(config.BasePath, path)
		if err != nil {
			return nil
		}

		suggestions = append(suggestions, relativePath)
		return nil
	})
	// If an error occurs during traversal, return an empty slice
	if err != nil {
		return nil
	}

	sort.Strings(suggestions)
	return suggestions
}
