// +build go1.16

package goinit

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed template/*
var tpl embed.FS

func run(wd string, params *templateParams, textTpl *template.Template) error {
	return fs.WalkDir(tpl, "template", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		pathToFile := strings.Replace(filepath.Dir(path)[len("template"):], "name", params.Name, 1)
		pathToFile = strings.Replace(pathToFile, "sub", params.SubName, 1)
		dir := wd + pathToFile
		fileName := strings.Replace(d.Name(), "name", params.Name, 1)
		fileName = strings.Replace(fileName, "sub", params.SubName, 1)

		rawTpl, err := tpl.ReadFile(path)
		if err != nil {
			return fmt.Errorf("tpl.ReadFile: %s, path: %s", err, path)
		}

		return create(dir, fileName, params, textTpl, rawTpl)
	})
}
