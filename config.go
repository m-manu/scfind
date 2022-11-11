package main

import (
	_ "embed"
	"encoding/json"
	"strings"
)

//go:embed config_allowed_file_extensions.txt
var allowedFileExtensionsRaw string

//go:embed config_allowed_file_names.txt
var allowedFileNamesRaw string

//go:embed config_ignored_directories.txt
var ignoredDirectoriesRaw string

//go:embed config_ignored_directories_with_peer_file_names.json
var ignoredDirectoriesWithPeerFileNamesRaw []byte

var allowedFileExtensions set[string]
var allowedFileNames set[string]
var ignoredDirectories set[string]
var ignoredDirectoriesWithPeerFileNames map[string][]string

func toLookupMap(ldEntries string) set[string] {
	entries := strings.Split(ldEntries, "\n")
	s := newSet[string](len(entries))
	for _, entry := range entries {
		if strings.TrimSpace(entry) == "" {
			continue
		}
		s.add(entry)
	}
	return s
}

func init() {
	allowedFileExtensions = toLookupMap(allowedFileExtensionsRaw)
	allowedFileNames = toLookupMap(allowedFileNamesRaw)
	ignoredDirectories = toLookupMap(ignoredDirectoriesRaw)
	_ = json.Unmarshal(ignoredDirectoriesWithPeerFileNamesRaw, &ignoredDirectoriesWithPeerFileNames)
}
