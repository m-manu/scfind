# scfind

[![build](https://github.com/m-manu/scfind/actions/workflows/build.yml/badge.svg)](https://github.com/m-manu/scfind/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/m-manu/scfind)](https://goreportcard.com/report/github.com/m-manu/scfind)
[![Go Reference](https://pkg.go.dev/badge/github.com/m-manu/scfind.svg)](https://pkg.go.dev/github.com/m-manu/scfind)
[![License](https://img.shields.io/badge/License-Apache%202-blue.svg)](./LICENSE)

**scfind**, which stands for '**s**ource **c**ode **find**', is a replacement for `find` command for _source code_
files.

## How to install?

```shell
go install github.com/m-manu/scfind@latest
```

## How to use?

```shell
scfind <directory-path> 
```

Examples:

```shell
scfind ~/Programming
```

```shell
scfind . | xargs grep --color "LinkedHashSet"
```
