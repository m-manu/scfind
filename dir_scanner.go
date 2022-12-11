package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var prefix string

var applyFunc func(string)

func walkFunction(path string, de fs.DirEntry, err error) error {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "skipping \"%s\": %+v\n", path, err)
		return nil
	}
	// If the file/directory is in excluded files list, ignore it
	if ignoredDirectories.contains(de.Name()) && de.IsDir() {
		return filepath.SkipDir
	}
	if peerFiles, exists := ignoredDirectoriesWithPeerFileNames[de.Name()]; exists && de.IsDir() {
		for _, peerFile := range peerFiles {
			peerFilePath := filepath.Join(filepath.Dir(path), peerFile)
			if doesFileExist(peerFilePath) { // if such a peer file exists
				return filepath.SkipDir
			}
		}
	}
	// Ignore dot files (Mac) and non-regular files
	if strings.HasPrefix(de.Name(), "._") || !de.Type().IsRegular() {
		return nil
	}
	if allowedFileExtensions.contains(getFileExt(de.Name())) || allowedFileNames.contains(de.Name()) {
		applyFunc(path)
	}
	return nil
}

func setPrefix(d string) {
	if d == "." {
		prefix = "./"
	}
}

func scanDirectory(dir string, f func(string)) error {
	applyFunc = f
	setPrefix(dir)
	wdErr := filepath.WalkDir(dir, walkFunction)
	return wdErr
}
