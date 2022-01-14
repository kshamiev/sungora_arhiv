package test

import (
	"testing"

	"sungora/lib/tpl"

	"github.com/shopspring/decimal"
)

func TestTplStorage(t *testing.T) {
	cfg, ctx := GetEnv()

	task := tpl.NewTaskTemplateParse(cfg.App.DirWww)
	if err := task.Action(ctx); err != nil {
		t.Fatal(err)
	}

	goods := Goods{
		{ID: 37, Name: "Item 10", Price: decimal.NewFromFloat(23.76)},
		{ID: 49, Name: "Item 2", Price: decimal.NewFromFloat(87.42)},
		{ID: 54, Name: "Item 30", Price: decimal.NewFromFloat(38.23)},
	}

	variable := map[string]interface{}{
		"Title": "Fantik",
		"Goods": goods,
	}

	for _, i := range tpl.GetTplIndex() {
		t.Log(i)
		if _, err := tpl.ExecuteStorage(i, variable); err != nil {
			t.Fatal(err)
		}
	}
}

func TestTpl(t *testing.T) {
	goods := Goods{
		{ID: 23, Name: "Item 1", Price: decimal.NewFromFloat(45.76)},
		{ID: 34, Name: "Item 2", Price: decimal.NewFromFloat(12.42)},
		{ID: 45, Name: "Item 3", Price: decimal.NewFromFloat(74.23)},
	}

	variable := map[string]interface{}{
		"Title": "Funtik",
		"Goods": goods,
	}

	_, err := tpl.ExecuteString(testTpl, nil, variable)
	if err != nil {
		t.Fatal(err)
	}
}

type Good struct {
	ID     uint64
	Name   string
	Price  decimal.Decimal
	Method Method
}

type Method struct{}

func (m *Method) Call() string {
	return "object method"
}

type Goods []Good

// language=html
const testTpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>

{{TplTest .Title}}
<table cellspacing="2" cellpadding="2">
	{{range .Goods}}
	{{if eq .Name "Item 2"}}
	<tr bgcolor="#f0fff0">
	{{else}}
	<tr bgcolor="#fff0f5">
	{{end}}
		<td>{{.ID}}</td>
		<td>{{.Name}}</td>
		<td>{{.Price}}</td>
		<td>{{.Method.Call}}</td>
	</tr>
	{{end}}
</table>

</body>
</html>
`
