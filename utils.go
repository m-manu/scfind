package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func getFileExt(fileName string) string {
	fileExt := strings.ReplaceAll(
		strings.ToLower(
			filepath.Ext(fileName),
		),
		".", "")
	return fileExt
}

func checkDirectoryIsReadable(path string) error {
	fileInfo, statErr := os.Stat(path)
	if statErr != nil {
		return statErr
	}
	if !fileInfo.IsDir() {
		return errors.New("not a directory")
	}
	return nil
}

func doesFileExist(path string) bool {
	info, err := os.Stat(path)
	return err == nil &&
		info.Mode().IsRegular()
}
