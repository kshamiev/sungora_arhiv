package tpl

import (
	"text/template"
	"time"
)

var tplStoreInfo = map[string]time.Time{}
var tplStore = map[string]*template.Template{}

var functions = map[string]interface{}{
	"TplTest": TplTest,
}

func TplTest(name string) string {
	return "<H1>" + name + "</H1>"
}
