# scfind

[![Build and test](https://github.com/m-manu/scfind/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/m-manu/scfind/actions/workflows/build-and-test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/m-manu/scfind)](https://goreportcard.com/report/github.com/m-manu/scfind)
[![Go Reference](https://pkg.go.dev/badge/github.com/m-manu/scfind.svg)](https://pkg.go.dev/github.com/m-manu/scfind)
[![License](https://img.shields.io/badge/License-Apache%202-blue.svg)](./LICENSE)

## Why?

`find` command in unix-like OSes is great. But, it is not helpful in searching for source code files or in searching
inside source code files (when used along with `grep`). So, we need a specialized command.

## What?

**scfind**, which stands for '**s**ource **c**ode **find**', is a replacement for `find` command for _source code_
files. It's ultra light and very fast.

## How to install?

1. Install Go version at least **1.18**
    * See: [Go installation instructions](https://go.dev/doc/install)
2. Run command:
   ```bash
   go install github.com/m-manu/scfind@latest
   ```
3. Add following line in your `.bashrc`/`.zshrc` file:
   ```bash
   export PATH="$PATH:$HOME/go/bin"
   ```

## How to use?

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

`scfind` command traverses file tree with source code awareness in following ways:

1. Scans for files with known source code and configuration file extensions (case insensitive)
    * e.g.`.java`, `.go`, `.py`, `.yml` etc.
    * see [full list](config/allowed_file_extensions.txt)
2. Scans for files with certain names (case sensitive)
    * e.g. `postinst`, `Dockerfile` etc.
    * see [full list](config/allowed_file_names.txt)
3. Skips scanning certain directories (case sensitive)
    * e.g. `.git`, `.idea`, `.gradle` etc.
    * see [full list](config/ignored_directories.txt)
4. Skips scanning certain directories with specific peer files (case sensitive)
    * e.g. skip `build` sub-directory when `build.gradle` exists in the same directory etc.
    * see [full list](config/ignored_directories_with_peer_file_names.json)
