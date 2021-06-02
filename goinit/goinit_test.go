package goinit

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	wd, err := os.Getwd()
	assert.NoError(t, err)

	repo := "path/to/repo"
	name := "foo"
	subName := "bar"
	appDir := "internal"
	ttpl := template.New("goinit")

	testCases := map[string]struct {
		goinit *goinit
		paths  map[string]string
	}{
		"main": {
			goinit: &goinit{
				params: &templateParams{
					Repo:      repo,
					Name:      name,
					UpperName: strings.Title(name),
					CmdDir:    appDir,
				},
				tpl: ttpl,
			},
			paths: map[string]string{
				"template/cmd/name/name.go.txt":                     fmt.Sprintf("cmd/%s/%s.go", name, name),
				"template/internal/cmd/name/name.go.txt":            fmt.Sprintf("%s/cmd/%s/%s.go", appDir, name, name),
				"template/internal/cmd/name/options/options.go.txt": fmt.Sprintf("%s/cmd/%s/options/options.go", appDir, name),
			},
		},
		"sub": {
			goinit: &goinit{
				params: &templateParams{
					Repo:         repo,
					Name:         name,
					UpperName:    strings.Title(name),
					SubName:      subName,
					UpperSubName: strings.Title(subName),
					CmdDir:       appDir,
				},
				tpl: ttpl,
			},
			paths: map[string]string{
				"template/cmd/name/name.go.txt":                     fmt.Sprintf("cmd/%s/%s.go", name, name),
				"template/internal/cmd/name/name.go.txt":            fmt.Sprintf("%s/cmd/%s/%s.go", appDir, name, name),
				"template/internal/cmd/name/options/options.go.txt": fmt.Sprintf("%s/cmd/%s/options/options.go", appDir, name),
				"template/internal/cmd/name/sub/sub.go.txt":         fmt.Sprintf("%s/cmd/%s/%s/%s.go", appDir, name, subName, subName),
			},
		},
	}

	for testCase, tt := range testCases {
		t.Run(testCase, func(t *testing.T) {
			tempDir := t.TempDir()
			assert.NoError(t, tt.goinit.run(tempDir))

			for tplPath, path := range tt.paths {
				rawTpl, err := readFile(filepath.FromSlash(wd + "/" + tplPath))
				assert.NoError(t, err)
				tpl, err := tt.goinit.tpl.Parse(rawTpl)
				assert.NoError(t, err)
				wantCode := &bytes.Buffer{}
				assert.NoError(t, tpl.Execute(wantCode, tt.goinit.params))

				gotCode, err := ioutil.ReadFile(filepath.FromSlash(tempDir + "/" + path))
				assert.NoError(t, err)
				// compare as string for debug
				assert.Equal(t, wantCode.String(), string(gotCode))
			}
		})
	}
}
