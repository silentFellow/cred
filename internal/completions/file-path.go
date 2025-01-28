package completions

import (
	"os"
	"path/filepath"
)

func GetFilePathSuggestions(path, basePath string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	if len(files) == 0 {
		return []string{}
	}

	suggestions := make([]string, 0)

	// For each file/directory in the current path
	for _, file := range files {
		// Construct the full relative path (e.g., "hi/email.gpg")
		fullPath := filepath.Join(path, file.Name())

		// If it's a directory, append it with '/' at the end
		if file.IsDir() {
			// Recursively get suggestions from the subdirectory
			subSuggestions := GetFilePathSuggestions(fullPath, basePath)

			// Append the subdirectory paths, using full path
			for _, subSuggestion := range subSuggestions {
				suggestions = append(suggestions, subSuggestion)
			}
		} else {
			// If it's a file, add the full relative path to the suggestions
			relativePath, _ := filepath.Rel(basePath, fullPath)
			suggestions = append(suggestions, relativePath)
		}
	}

	return suggestions
}
