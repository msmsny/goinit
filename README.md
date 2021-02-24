# goinit

[![Go Reference](https://pkg.go.dev/badge/github.com/msmsny/goinit.svg)](https://pkg.go.dev/github.com/msmsny/goinit)
[![Go Report Card](https://goreportcard.com/badge/github.com/msmsny/goinit)](https://goreportcard.com/report/github.com/msmsny/goinit)

Go project init tool

## Install

```bash
$ go get github.com/msmsny/goinit
```

## Usage

```bash
$ goinit -help
Usage of /path/to/bin/goinit:
  -app-dir string
    	application code directory (default "internal")
  -name string
    	main command name(required)
  -repo string
    	path to repo, e.g.: github.com/msmsny/goinit (default "github.com/msmsny/goinit")
  -sub string
    	sub command name
```

```bash
$ goinit -name foo
```

```bash
$ tree
.
├── cmd
│   └── foo
│       └── foo.go
└── internal
    ├── cmd
    │   └── foo
    │       ├── foo.go
    │       └── options
    │           └── options.go
    └── foo
        └── foo.go
```

```bash
$ goinit -name foo -sub bar
```

```bash
$ tree
.
├── cmd
│   └── foo
│       └── foo.go
└── internal
    ├── cmd
    │   └── foo
    │       ├── bar
    │       │   └── bar.go
    │       ├── foo.go
    │       └── options
    │           └── options.go
    └── foo
        └── foo.go
```
