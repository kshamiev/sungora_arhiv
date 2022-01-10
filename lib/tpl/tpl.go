package tpl

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"text/template"

	"sungora/lib/errs"
)

// GetTplIndex индексы подготовленных шаблонов
func GetTplIndex() []string {
	s := make([]string, 0, len(tplStore))
	for i := range tplStore {
		s = append(s, i)
	}
	return s
}

// ExecuteStorage сборка контента из подготовленного шаблона
func ExecuteStorage(viewIndex string, variables map[string]interface{}) (ret bytes.Buffer, err error) {
	tpl, ok := tplStore[viewIndex]
	if !ok {
		return ret, errs.NewNotFound(errors.New("not found tpl: " + viewIndex))
	}
	err = tpl.Execute(&ret, variables)
	return
}

// ExecuteFile компиляция html шаблона из указанного файла и сборка контента
func ExecuteFile(viewPath string, funcs, variables map[string]interface{}) (ret bytes.Buffer, err error) {
	if _, err = os.Stat(viewPath); err != nil {
		return ret, errs.NewBadRequest(err)
	}

	var tpl *template.Template

	if funcs == nil {
		funcs = functions
	}

	if tpl, err = template.New(filepath.Base(viewPath)).Funcs(funcs).ParseFiles(viewPath); err != nil {
		return ret, errs.NewBadRequest(err)
	}

	if err = tpl.Execute(&ret, variables); err != nil {
		return ret, errs.NewBadRequest(err)
	}

	return
}

// ExecuteString компиляция html шаблона переданного в строке и сборка контента
func ExecuteString(view string, funcs, variables map[string]interface{}) (ret bytes.Buffer, err error) {
	const nameTpl = "default"
	var tpl *template.Template

	if funcs == nil {
		funcs = functions
	}

	if tpl, err = template.New(nameTpl).Funcs(funcs).Parse(view); err != nil {
		return ret, errs.NewBadRequest(err)
	}

	if err = tpl.Execute(&ret, variables); err != nil {
		return ret, errs.NewBadRequest(err)
	}

	return
}
