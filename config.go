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

type stringLookup map[string]bool

var allowedFileExtensions stringLookup
var allowedFileNames stringLookup
var ignoredDirectories stringLookup
var ignoredDirectoriesWithPeerFileNames map[string][]string

func toLookupMap(ldEntries string) stringLookup {
	entries := strings.Split(ldEntries, "\n")
	m := make(stringLookup, len(entries))
	for _, s := range entries {
		if strings.TrimSpace(s) == "" {
			continue
		}
		m[s] = true
	}
	return m
}

func init() {
	allowedFileExtensions = toLookupMap(allowedFileExtensionsRaw)
	allowedFileNames = toLookupMap(allowedFileNamesRaw)
	ignoredDirectories = toLookupMap(ignoredDirectoriesRaw)
	_ = json.Unmarshal(ignoredDirectoriesWithPeerFileNamesRaw, &ignoredDirectoriesWithPeerFileNames)
}
