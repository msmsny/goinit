// +build go1.16

package goinit

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed template/*
var tpl embed.FS

func (g *goinit) run(wd string) error {
	return fs.WalkDir(tpl, "template", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if g.params.SubName == "" && strings.Contains(path, filepath.FromSlash("/sub")) {
			return nil
		}

		pathToFile := strings.Replace(filepath.Dir(path)[len("template"):], "name", g.params.Name, 1)
		pathToFile = strings.Replace(pathToFile, "sub", g.params.SubName, 1)
		dir := wd + pathToFile
		fileName := strings.Replace(d.Name(), "name", g.params.Name, 1)
		fileName = strings.Replace(fileName, "sub", g.params.SubName, 1)

		rawTpl, err := tpl.ReadFile(path)
		if err != nil {
			return fmt.Errorf("tpl.ReadFile: %s, path: %s", err, path)
		}

		return g.generate(dir, fileName, rawTpl)
	})
}
