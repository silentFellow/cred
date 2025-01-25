package fscopy

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Copy(src, dest string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		fmt.Println("Failed to stat source file, ", err)
		return err
	}

	if srcInfo.IsDir() {
		return CopyDirectory(src, dest)
	}

	return CopyFile(src, dest)
}

func CopyDirectory(src, dest string) error {
	if err := os.MkdirAll(dest, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create destination directory: %v", err)
	}

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("Failed to access %q file: %w", path, err)
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path of file %v: %w", path, err)
		}
		detPath := filepath.Join(dest, relPath)

		if info.IsDir() {
			if err := os.MkdirAll(detPath, info.Mode()); err != nil {
				return fmt.Errorf("failed to create directory %v: %w", detPath, err)
			}
		} else {
			if err := CopyFile(path, detPath); err != nil {
				return fmt.Errorf("failed to copy file from %v to %v: %w", path, detPath, err)
			}
		}

		return nil
	})
}

func CopyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Failed to open source file: %w", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("Failed to create destination file: %w", err)
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return fmt.Errorf("Failed to copy file: %w", err)
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("Failed to stat source file: %w", err)
	}

	if err := os.Chmod(dest, srcInfo.Mode()); err != nil {
		return fmt.Errorf("Failed to change permission of the destination file: %w", err)
	}

	return nil
}
