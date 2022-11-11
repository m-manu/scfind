package main

import (
	"os"
	"strings"
	"testing"
)

var pathsToTest = []string{
	"config_allowed_file_extensions.txt",
	"config_allowed_file_names.txt",
	"config_ignored_directories.txt",
}

func TestConfigFiles(t *testing.T) {
	for _, path := range pathsToTest {
		if !doesFileExist(path) {
			t.Errorf("File doesn't exist: %s", path)
		}
		contentsBytes, _ := os.ReadFile(path)
		contents := string(contentsBytes)
		entries := newSet[string](100)
		lines := strings.Split(contents, "\n")
		for lineNumber, lineText := range lines {
			if strings.TrimSpace(lineText) == "" {
				t.Errorf("Issue in file %s: Line %d is empty", path, lineNumber+1)
			}
			if entries.contains(lineText) {
				t.Errorf("Issue in file %s at line %d: Entry \"%s\" is repeated", path, lineNumber+1, lineText)
			}
			entries.add(lineText)
		}
	}
}
