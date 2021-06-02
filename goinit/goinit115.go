// +build !go1.16

package goinit

//go:generate statik -src template

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/msmsny/goinit/goinit/statik"
	statikfs "github.com/rakyll/statik/fs"
)

func (g *goinit) run(wd string) error {
	httpfs, err := statikfs.New()
	if err != nil {
		return fmt.Errorf("statikfs.New: %s", err)
	}

	return statikfs.Walk(httpfs, "/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("statikfs.Walk: %s", err)
		}
		if info.IsDir() {
			return nil
		}
		if g.params.SubName == "" && strings.Contains(path, filepath.FromSlash("/sub")) {
			return nil
		}

		pathToFile := strings.Replace(filepath.Dir(path), "name", g.params.Name, 1)
		pathToFile = strings.Replace(pathToFile, "sub", g.params.SubName, 1)
		pathToFile = strings.Replace(pathToFile, filepath.FromSlash("internal/cmd"), filepath.FromSlash(g.params.CmdDir), 1)
		dir := wd + pathToFile
		fileName := strings.Replace(info.Name(), "name", g.params.Name, 1)
		fileName = strings.Replace(fileName, "sub", g.params.SubName, 1)

		file, err := httpfs.Open(path)
		if err != nil {
			return fmt.Errorf("httpfs.Open: %s", err)
		}
		rawTpl, err := ioutil.ReadAll(file)
		if err != nil {
			return fmt.Errorf("ioutil.ReadAll: %s", err)
		}

		return g.generate(dir, fileName, rawTpl)
	})
}

// readFile reads file from path and replaces CRLF for windows
func readFile(path string) (string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadFile: %s", err)
	}

	return strings.Replace(string(file), "\r\n", "\n", -1), nil
}
