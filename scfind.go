/*
scfind, which stands for 'source code find', is a replacement for 'find' command for source code files
*/
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	exitCodeInvalidNumArgs = iota + 1
	exitCodeScanFailed
	exitCodeInputDirectoryNotReadable
)

func getFileExt(fileName string) string {
	fileExt := strings.ReplaceAll(
		strings.ToLower(
			filepath.Ext(fileName),
		),
		".", "")
	return fileExt
}

func walkFunction(path string, de fs.DirEntry, err error) error {
	if err != nil {
		fmt.Fprintf(os.Stderr, "skipping \"%s\": %+v\n", path, err)
	}
	// If the file/directory is in excluded files list, ignore it
	if ignoredDirectories[de.Name()] && de.IsDir() {
		return filepath.SkipDir
	}
	if peerFiles, exists := ignoredDirectoriesWithPeerFileNames[de.Name()]; exists && de.IsDir() {
		for _, peerFile := range peerFiles {
			peerFilePath := filepath.Join(filepath.Dir(path), peerFile)
			if _, fErr := os.Stat(peerFilePath); fErr == nil { // if such a peer file exists
				return filepath.SkipDir
			}
		}
	}
	// Ignore dot files (Mac) and non-regular files
	if strings.HasPrefix(de.Name(), "._") || !de.Type().IsRegular() {
		return nil
	}
	if allowedFileExtensions[getFileExt(de.Name())] || allowedFileNames[de.Name()] {
		printFileName(path)
	}
	return nil
}

func printFileName(path string) {
	fmt.Printf("%s\n", path)
}
func isReadableDirectory(path string) bool {
	fileInfo, statErr := os.Stat(path)
	if statErr != nil {
		return false
	}
	return fileInfo.IsDir()
}

func readDirectory(directory string) string {
	if !isReadableDirectory(directory) {
		fmt.Fprintf(os.Stderr, "error: input \"%v\" isn't a readable directory\n", directory)
		flag.Usage()
		os.Exit(exitCodeInputDirectoryNotReadable)
	}
	directory = filepath.Clean(directory)
	return directory
}

const helpMessage = `scfind is a "find command for source code files"

Usage: 
	scfind <directory-path>

For more details: https://github.com/m-manu/scfind`

func main() {
	flag.Usage = func() {
		fmt.Println(helpMessage)
	}
	flag.Parse()
	if flag.NArg() > 1 || flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "error: pass only one argument (directory path)\n")
		os.Exit(exitCodeInvalidNumArgs)
		return
	}
	directoryToScan := readDirectory(flag.Arg(0))
	err := filepath.WalkDir(directoryToScan, walkFunction)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: couldn't read current directory: %+v", err)
		os.Exit(exitCodeScanFailed)
	}
}
