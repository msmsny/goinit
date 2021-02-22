package main

import (
	"fmt"
	"os"

	cmd "{{.Repo}}/pkg/cmd/{{.Name}}"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
