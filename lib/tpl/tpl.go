package tpl

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var tplStore = map[string]*template.Template{}
var Functions = map[string]interface{}{}

func Init(dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "temp" {
			return filepath.SkipDir
		}
		if info.IsDir() || filepath.Ext(path) != ".html" {
			return nil
		}
		if err = parseFiles(dir, path); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
}

// компиляция html шаблонов
func parseFiles(rootDir, viewPath string) error {
	data, err := ioutil.ReadFile(viewPath)
	if err != nil {
		return err
	}
	index := strings.ReplaceAll(viewPath, rootDir+"/", "")

	tpl, err := template.New(index).Funcs(Functions).Parse(string(data))
	if err != nil {
		return err
	}

	tplStore[index] = tpl
	return nil
}

// сборка контента из подготовленного шаблона
func Execute(index string, variables map[string]interface{}) (ret bytes.Buffer, err error) {
	tpl, ok := tplStore[index]
	if !ok {
		return ret, errors.New("not found tpl: " + index)
	}
	err = tpl.Execute(&ret, variables)
	return
}

// индексы подготовленных шаблонов
func GetTplIndex() []string {
	s := make([]string, 0, len(tplStore))
	for i := range tplStore {
		s = append(s, i)
	}
	return s
}
