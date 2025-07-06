/*
*
scfind, which stands for 'source code find', is a replacement for 'find' command for source code files
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	exitSuccess = iota
	exitCodeInvalidNumArgs
	exitCodeScanFailed
	exitCodeInputDirectoryNotReadable
	exitCodeSymLinkEvalFailed
)

const helpMessage = `scfind is a 'find' command for source code files

Usage: 
	scfind $DIRECTORY_PATH

where,
	$DIRECTORY_PATH is path to a readable directory that
	you want to scan for source code files

For more details: https://github.com/m-manu/scfind`

func showErrorMessageAndExit(msg string, exitCode int) {
	_, _ = fmt.Fprintf(os.Stderr, "%s\n%s\n", msg, "Run `scfind -h` for usage")
	os.Exit(exitCode)
}

func printFileName(path string) {
	fmt.Printf("%s%s\n", prefix, path)
}

var flags struct {
	isHelp bool
}

func setupAndParseFlags() (isHelp bool, arg0 string) {
	flag.BoolVar(&flags.isHelp, "h", false, "print help with version information and exit")
	flag.Parse()
	if flags.isHelp {
		return true, ""
	}
	if flag.NArg() > 1 || flag.NArg() < 1 {
		showErrorMessageAndExit(
			fmt.Sprintf("error: accepts only one argument (directory path)"), exitCodeInvalidNumArgs)
		return
	}
	return false, flag.Arg(0)
}

func main() {
	isHelp, directory := setupAndParseFlags()
	if isHelp {
		fmt.Println(helpMessage)
		os.Exit(exitSuccess)
	}
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
	wdErr := scanDirectory(realDirectory, printFileName)
	if wdErr != nil {
		showErrorMessageAndExit(fmt.Sprintf("error: couldn't read directory: %+v", wdErr), exitCodeScanFailed)
	}
}
