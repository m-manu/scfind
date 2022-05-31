# Source Code Find (scfind)

[![build](https://github.com/m-manu/scfind/actions/workflows/build.yml/badge.svg)](https://github.com/m-manu/scfind/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/m-manu/scfind)](https://goreportcard.com/report/github.com/m-manu/scfind)
[![Go Reference](https://pkg.go.dev/badge/github.com/m-manu/scfind.svg)](https://pkg.go.dev/github.com/m-manu/scfind)
[![License](https://img.shields.io/badge/License-Apache%202-blue.svg)](./LICENSE)

**scfind**, which stands for '**s**ource **c**ode **find**', is a replacement for `find` command for _source code_
files. It's ultra light and very fast.

## How to install?

1. Install Go version at least **1.16**
    * See: [Go installation instructions](https://go.dev/doc/install)
2. Run command:
   ```bash
   go install github.com/m-manu/scfind@latest
   ```
3. Add following line in your `.bashrc`/`.zshrc` file:
   ```bash
   export PATH="$PATH:$HOME/go/bin"
   ```

## Usage

Running `scfind -h` shows this help message:

```text
scfind is a 'find' command for source code files

Usage: 
	scfind DIRECTORY_PATH

where,
	DIRECTORY_PATH is path to a readable directory that
	you want to scan for source code files

For more details: https://github.com/m-manu/scfind
```

### Examples

```shell
scfind ~/Programming
```

```shell
scfind . | xargs grep --color "LinkedHashSet"
```

## How does this work?

scfind traverses file tree with source code awareness in following ways:

1. Skips scanning certain directories (case sensitive)
    * e.g. `.git`, `.idea` etc.
    * see [full list](./config_ignored_directories.txt)
2. Skips scanning certain directories with specific peer files (case sensitive)
    * e.g. skip `build` sub-directory when `build.gradle` exists in the same directory etc.
    * see [full list](./config_ignored_directories_with_peer_file_names.json)
3. Lists out files only with known source code file extensions (case insensitive)
    * e.g.`.java`, `.go`, `.py` etc.
    * see [full list](./config_allowed_file_extensions.txt)
4. Lists out files with certain names (case sensitive)
    * e.g. `postinst`, `Dockerfile` etc.
    * see [full list](./config_allowed_file_names.txt)
