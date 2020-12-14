package tpl

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestInit(t *testing.T) {
	Init("../../template")
	for _, i := range GetTplIndex() {
		if _, err := Execute(i, nil); err != nil {
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

	data, err := ParseText(testTpl, Functions, variable)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data.String())
}

type Good struct {
	ID    uint64
	Name  string
	Price decimal.Decimal
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

{{Test .Title}}
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
	</tr>
	{{end}}
</table>

</body>
</html>
`
