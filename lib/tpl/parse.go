package tpl

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// компиляция html шаблона из указанного файла и сборка контента
func ParseFile(viewPath string, functions, variables map[string]interface{}) (ret bytes.Buffer, err error) {
	if _, err = os.Stat(viewPath); err != nil {
		return
	}

	var tpl *template.Template

	if tpl, err = template.New(filepath.Base(viewPath)).Funcs(functions).ParseFiles(viewPath); err != nil {
		return
	}

	if err = tpl.Execute(&ret, variables); err != nil {
		return
	}

	return
}

// компиляция html шаблона переданного в строке и сборка контента
func ParseText(view string, functions, variables map[string]interface{}) (ret bytes.Buffer, err error) {
	const nameTpl = "default"
	var tpl *template.Template

	if tpl, err = template.New(nameTpl).Funcs(functions).Parse(view); err != nil {
		return
	}

	if err = tpl.Execute(&ret, variables); err != nil {
		return
	}

	return
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
