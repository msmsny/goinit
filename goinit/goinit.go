package goinit

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func Run() error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd: %s", err)
	}

	pathToRepo := filepath.Dir(filepath.Join(wd, "..", ".."))
	repo := flag.String("repo", wd[len(pathToRepo+"/"):], "path to repo, e.g.: github.com/msmsny/goinit")
	name := flag.String("name", "", "main command name(required)")
	subName := flag.String("sub", "", "sub command name")
	cmdDir := flag.String("cmd-dir", filepath.FromSlash("pkg/cmd"), "cobra command code directory")
	flag.Parse()

	if *name == "" {
		flag.Usage()
		os.Exit(1)
	}

	goinit := &goinit{
		params: &templateParams{
			Repo:         *repo,
			Name:         *name,
			UpperName:    strings.Title(*name),
			SubName:      *subName,
			UpperSubName: strings.Title(*subName),
			CmdDir:       *cmdDir,
		},
		tpl: template.New("goinit"),
	}

	return goinit.run(wd)
}

type templateParams struct {
	Repo         string
	Name         string
	UpperName    string
	SubName      string
	UpperSubName string
	CmdDir       string
}

type goinit struct {
	params *templateParams
	tpl    *template.Template
}

func (g *goinit) generate(dir, fileName string, rawTpl []byte) error {
	// create directory if not exist
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("os.MkdirAll: %s", err)
			}
		} else {
			return fmt.Errorf("os.Stat: %s", err)
		}
	}

	// embed template parameters
	t, err := g.tpl.Parse(string(rawTpl))
	if err != nil {
		return fmt.Errorf("template.Parse: %s", err)
	}
	contents := &bytes.Buffer{}
	if err := t.Execute(contents, g.params); err != nil {
		return fmt.Errorf("template.Execute: %s, params: %#v", err, g.params)
	}

	// create file
	filePath := fmt.Sprintf(filepath.FromSlash("%s/%s"), dir, fileName[:len(fileName)-len(".txt")])
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("os.Create: %s, path: %s", err, filePath)
	}
	defer file.Close()
	if _, err := file.Write(contents.Bytes()); err != nil {
		return fmt.Errorf("file.Write: %s", err)
	}

	return nil
}
