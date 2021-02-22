// +build go1.16

package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed template/*
var tpl embed.FS

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd: %s", err)
	}

	pathToRepo := filepath.Dir(filepath.Join(wd, "..", ".."))
	repo := flag.String("repo", wd[len(pathToRepo+"/"):], "path to repo, e.g.: github.com/msmsny/goinit")
	name := flag.String("name", "", "main command name(required)")
	subName := flag.String("sub", "sub", "sub command name")
	flag.Parse()

	if *name == "" {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	params := &templateParams{
		Repo:         *repo,
		Name:         *name,
		UpperName:    strings.Title(*name),
		SubName:      *subName,
		UpperSubName: strings.Title(*subName),
	}
	textTpl := template.New("goinit")

	fs.WalkDir(tpl, "template", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		pathToFile := strings.Replace(filepath.Dir(path)[len("template"):], "name", *name, 1)
		pathToFile = strings.Replace(pathToFile, "sub", *subName, 1)
		dir := wd + pathToFile
		fileName := strings.Replace(d.Name(), "name", *name, 1)
		fileName = strings.Replace(fileName, "sub", *subName, 1)

		// create directory if not exist
		if _, err := os.Stat(dir); err != nil {
			if os.IsNotExist(err) {
				os.MkdirAll(dir, 0755)
			} else {
				log.Fatalf("os.Stat: %s", err)
			}
		}

		// embed template parameters
		rawTpl, err := tpl.ReadFile(path)
		if err != nil {
			log.Fatalf("read file %s failed: %s", path, err)
		}
		t, err := textTpl.Parse(string(rawTpl))
		if err != nil {
			log.Fatalf("template.Parse: %s", err)
		}
		contents := &bytes.Buffer{}
		if err := t.Execute(contents, params); err != nil {
			log.Fatalf("template.Execute: %s, params: %#v", err, params)
		}

		// create file
		filePath := fmt.Sprintf("%s/%s", dir, fileName[:len(fileName)-len(".tpl")])
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("os.Create: %s, path: %s", err, filePath)
		}
		defer file.Close()
		file.Write(contents.Bytes())

		return nil
	})
}

type templateParams struct {
	Repo         string
	Name         string
	UpperName    string
	SubName      string
	UpperSubName string
}
