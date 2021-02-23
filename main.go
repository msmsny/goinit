package main

import (
	"fmt"
	"os"

	"github.com/msmsny/goinit/pkg/goinit"
)

func main() {
	if err := goinit.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
