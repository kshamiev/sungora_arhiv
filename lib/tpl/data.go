package tpl

import (
	"text/template"
)

var tplStore = map[string]*template.Template{}
var Functions = map[string]interface{}{
	"Test": Test,
}

func Test(name string) string {
	return "<H1>" + name + "</H1>"
}
