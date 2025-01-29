package paths

import (
	"os"
	"strings"
)

func CheckPathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

func GetPathType(path string) string {
	info, _ := os.Stat(path)

	fileType := "file"
	if info.IsDir() {
		fileType = "directory"
	}

	return fileType
}

func CreatePath(path string) (*os.File, error) {
	pathFields := strings.Split(path, "/")

	var dir strings.Builder
	for i := 0; i < len(pathFields)-1; i++ {
		dir.WriteString(pathFields[i])
		dir.WriteString("/")
	}

	if err := os.MkdirAll(dir.String(), 0700); err != nil {
		return nil, err
	}

	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func BuildPath(pathSequence ...string) string {
	var path strings.Builder

	for i, currentPath := range pathSequence {
		path.WriteString(currentPath)

		if i != len(pathSequence)-1 && !strings.HasSuffix(currentPath, "/") {
			path.WriteString("/")
		}
	}

	return path.String()
}
