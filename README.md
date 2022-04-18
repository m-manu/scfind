# scfind

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
