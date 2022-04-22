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
	exitCodeSymLinkEvalFailed
)

func walkFunction(path string, de fs.DirEntry, err error) error {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "skipping \"%s\": %+v\n", path, err)
		return nil
	}
	// If the file/directory is in excluded files list, ignore it
	if ignoredDirectories[de.Name()] && de.IsDir() {
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
	if allowedFileExtensions[getFileExt(de.Name())] || allowedFileNames[de.Name()] {
		printFileName(path)
	}
	return nil
}

var prefix string

func setPrefix(d string) {
	if d == "." {
		prefix = "./"
	}
}

func printFileName(path string) {
	fmt.Printf("%s%s\n", prefix, path)
}

const helpMessage = `scfind is a 'find' command for source code files

Usage: 
	scfind DIRECTORY_PATH

where,
	DIRECTORY_PATH is path to a readable directory that
	you want to scan for source code files

For more details: https://github.com/m-manu/scfind`

func showErrorMessageAndExit(msg string, exitCode int) {
	_, _ = fmt.Fprintf(os.Stderr, "%s\n%s\n", msg, "Run `scfind -h` for usage")
	os.Exit(exitCode)
}

func main() {
	flag.Usage = func() {
		fmt.Println(helpMessage)
	}
	flag.Parse()
	if flag.NArg() > 1 || flag.NArg() < 1 {
		showErrorMessageAndExit(
			fmt.Sprintf("error: accepts only one argument (directory path)"), exitCodeInvalidNumArgs)
		return
	}
	directory := flag.Arg(0)
	dErr := checkDirectoryIsReadable(directory)
	if dErr != nil {
		showErrorMessageAndExit(
			fmt.Sprintf("error: input \"%v\" isn't a readable directory: %+v", directory, dErr),
			exitCodeInputDirectoryNotReadable,
		)
		return
	}
	realDirectory, slErr := filepath.EvalSymlinks(directory)
	if slErr != nil {
		showErrorMessageAndExit(fmt.Sprintf("error: unable to evaluate sym link: %+v", slErr),
			exitCodeSymLinkEvalFailed)
		return
	}
	setPrefix(realDirectory)
	wdErr := filepath.WalkDir(realDirectory, walkFunction)
	if wdErr != nil {
		showErrorMessageAndExit(fmt.Sprintf("error: couldn't read directory: %+v", wdErr), exitCodeScanFailed)
	}
}
