package fscopy

import (
	"io"
	"os"
	"path/filepath"
)

func Copy(src, dest string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if srcInfo.IsDir() {
		return CopyDirectory(src, dest)
	}

	return CopyFile(src, dest)
}

func CopyDirectory(src, dest string) error {
	if err := os.MkdirAll(dest, os.ModePerm); err != nil {
		return err
	}

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		detPath := filepath.Join(dest, relPath)

		if info.IsDir() {
			if err := os.MkdirAll(detPath, info.Mode()); err != nil {
				return err
			}
		} else {
			if err := CopyFile(path, detPath); err != nil {
				return err
			}
		}

		return nil
	})
}

func CopyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Check if `dest` is a directory and adjust the destination path
	destInfo, err := os.Stat(dest)
	if err != nil {
		return err
	} else if destInfo.IsDir() {
		dest = filepath.Join(dest, filepath.Base(src)) // Append filename
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return err
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.Chmod(dest, srcInfo.Mode()); err != nil {
		return err
	}

	return nil
}
