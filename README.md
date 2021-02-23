# goinit

Go project init tool

## Install

```bash
$ go get github.com/msmsny/goinit
```

## Usage

```bash
$ goinit -help
Usage of /path/to/bin/goinit:
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
└── pkg
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
└── pkg
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
