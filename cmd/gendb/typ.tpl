// nolint: lll // AFTER CODE GENERATED. DO NOT EDIT //
package typ

import (
	"github.com/shopspring/decimal"

	"sungora/lib/null"
	"sungora/lib/uuid"
)

// {{.schema.Table.Description.String}}
type {{.schema.Table.NameType}} struct {
{{range .schema.Columns}} {{.NameType}} {{.TypeFromGO}} `json:"{{.Name}}" db:"{{.Name}}"{{.TagFromGO}}` // {{.Description.String}}
{{end}} }

func (o *{{.schema.Table.NameType}}) Select() (query string, args []interface{}) {
	const SQL{{.schema.Table.NameType}}Select = "SELECT {{.schema.SqlColumnInsert}} FROM public.{{.schema.Table.Name}} WHERE id = $1"
	return SQL{{.schema.Table.NameType}}Select, []interface{}{
		o.Id,
	}
}

func (o *{{.schema.Table.NameType}}) Insert() (query string, args []interface{}) {
	const SQL{{.schema.Table.NameType}}Insert = "INSERT INTO public.{{.schema.Table.Name}} ({{.schema.SqlColumnInsert}}) VALUES ({{.schema.SqlColumnParams}})"
	return SQL{{.schema.Table.NameType}}Insert, []interface{}{
    {{range .schema.Columns}} o.{{.NameType}}, {{end}}
    }
}

func (o *{{.schema.Table.NameType}}) Update() (query string, args []interface{}) {
	const SQL{{.schema.Table.NameType}}Update = "UPDATE public.{{.schema.Table.Name}} SET {{.schema.SqlColumnUpdate}} WHERE id = $1"
	return SQL{{.schema.Table.NameType}}Update, []interface{}{
    {{range .schema.Columns}} o.{{.NameType}}, {{end}}
   	}
}

func (o *{{.schema.Table.NameType}}) Upsert() (query string, args []interface{}) {
	const SQL{{.schema.Table.NameType}}Upsert = "INSERT INTO public.{{.schema.Table.Name}} ({{.schema.SqlColumnInsert}}) VALUES ({{.schema.SqlColumnParams}}) ON CONFLICT (id) DO UPDATE SET {{.schema.SqlColumnUpdate}}"
	return SQL{{.schema.Table.NameType}}Upsert, []interface{}{
    {{range .schema.Columns}} o.{{.NameType}}, {{end}}
    }
}

func (o *{{.schema.Table.NameType}}) Delete() (query string, args []interface{}) {
	const SQL{{.schema.Table.NameType}}Delete = "DELETE FROM public.{{.schema.Table.Name}} WHERE id = $1"
	return SQL{{.schema.Table.NameType}}Delete, []interface{}{
		o.Id,
	}
}

// BEFORE CODE GENERATED. DO NOT EDIT //
