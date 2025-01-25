package paths

import (
	"os"
	"strings"

	gpgcrypt "github.com/silentFellow/cred-store/internal/gpg-crypt"
	"github.com/silentFellow/cred-store/internal/utils"
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


func AddToPath(path string, content string, copy bool) error {
	file, err := CreatePath(path)
	defer file.Close()

	if err != nil {
		return err
	}

	encrypted, err := gpgcrypt.Encrypt(content)
	if err != nil {
		return err
	}

	if _, err := file.WriteString(encrypted); err != nil {
		return err
	}

	if copy {
		// true since it only applicable for generate password
		if err := utils.CopyToClipboard(content, true); err != nil {
			return err
		}
	}

	return nil
}
