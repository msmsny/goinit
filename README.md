# goinit

Go project init tool

## Requirements

Go1.16+

## Install

```bash
$ go install github.com/msmsny/goinit
```

## Usage

```bash
$ goinit -help
Usage of /path/to/bin/goinit:
  -name string
    	main command name(required)
  -repo string
    	path to repo, e.g.: github.com/msmsny/goinit (default "github.com/msmsny/test")
  -sub string
    	sub command name (default "sub")
```

```bash
$ goinit -name foo -sub bar
```
