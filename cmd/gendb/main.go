// "github.com/volatiletech/null/v8"
// "github.com/volatiletech/sqlboiler/types"
package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/volatiletech/null"

	"sungora/lib/storage/stpg"
	"sungora/lib/tpl"
	"sungora/src/config"
)

func getConfig(fileConfig ...string) (cfg *Config, err error) {
	cfg = &Config{}
	for i := range fileConfig {
		if err = config.ConfigLoad(fileConfig[i], cfg); err == nil {
			return
		}
	}
	return
}

func nameConvert(name string) (nameType string) {
	for _, s := range strings.Split(name, "_") {
		nameType += strings.Title(s)
	}
	return nameType
}

func main() {
	flagConfigPath := flag.String("c", "conf/config.yaml", "used for set path to config file")
	flag.Parse()
	var err error

	generator := &Generator{}

	// Config загрузка конфигурации & Logger
	if generator.cfg, err = getConfig(*flagConfigPath); err != nil {
		log.Fatal(err)
	}

	// ConnectDB postgres
	if err = stpg.InitConnect(&generator.cfg.Postgresql); err != nil {
		log.Fatal(err)
	}

	generator.initTypeMappingDef()

	if err = generator.parsingDB(context.Background()); err != nil {
		log.Fatal(err)
	}

	if err = generator.saveFile(); err != nil {
		log.Fatal(err)
	}
}

type Generator struct {
	stpg.Storage
	cfg            *Config
	schema         []TableInfo
	TypeMappingDef map[string]map[string]string
}

func (gen *Generator) initTypeMappingDef() {
	gen.TypeMappingDef = map[string]map[string]string{
		"numeric": {
			"NO":  "decimal.Decimal",
			"YES": "decimal.Decimal",
			"Tag": ` example:"0.1"`,
		},
		"bool": {
			"NO":  "bool",
			"YES": "bool",
		},
		"text": {
			"NO":  "string",
			"YES": "null.String",
		},
		"varchar": {
			"NO":  "string",
			"YES": "null.String",
		},
		"timestamptz": {
			"NO":  "time.Time",
			"YES": "null.Time",
			"Tag": " example:\"2006-01-02T15:04:05Z\"",
		},
		"timestamp": {
			"NO":  "time.Time",
			"YES": "null.Time",
			"Tag": " example:\"2006-01-02T15:04:05Z\"",
		},
		"int": {
			"NO":  "int",
			"YES": "null.Int",
		},
		"int2": {
			"NO":  "int",
			"YES": "null.Int",
		},
		"int4": {
			"NO":  "int",
			"YES": "null.Int",
		},
		"int8": {
			"NO":  "int64",
			"YES": "null.Int64",
		},
		"float4": {
			"NO":  "float32",
			"YES": "null.Float32",
			"Tag": ` example:"0.1"`,
		},
		"float8": {
			"NO":  "float64",
			"YES": "null.Float64",
			"Tag": ` example:"0.1"`,
		},
	}
}

func (gen *Generator) parsingDB(ctx context.Context) error {
	var tables []*Table
	if err := gen.Query(ctx).Select(&tables, SQLTable, schema); err != nil {
		return err
	}
	gen.schema = make([]TableInfo, 0, len(tables))

	tableIgnore := make(map[string]bool)
	for i := range gen.cfg.Generate.TableIgnore {
		tableIgnore[gen.cfg.Generate.TableIgnore[i]] = true
	}

	for i := range tables {
		if _, ok := tableIgnore[tables[i].Name]; ok {
			continue
		}
		tables[i].NameType = nameConvert(tables[i].Name)
		tables[i].Description.String = strings.ReplaceAll(tables[i].Description.String, "\n", " ")
		tables[i].Description.String = strings.ReplaceAll(tables[i].Description.String, "\t", " ")

		columns := []*Column{}
		if err := gen.Query(ctx).Select(&columns, SQLColumn, schema, tables[i].Name); err != nil {
			return err
		}
		flagId := false
		columnsSqlInsert := make([]string, len(columns))
		columnsSqlParams := make([]string, len(columns))
		columnsSqlUpdate := make([]string, len(columns))
		for j := range columns {
			columns[j].NameType = nameConvert(columns[j].Name)
			columns[j].Description.String = strings.ReplaceAll(columns[j].Description.String, "\n", " ")
			columns[j].Description.String = strings.ReplaceAll(columns[j].Description.String, "\t", " ")
			columns[j].TypeFromGO, columns[j].TagFromGO = gen.typeConvert(columns[j])
			if columns[j].Name == "id" {
				flagId = true
			}
			columnsSqlUpdate[j] = columns[j].Name + " = " + "$" + strconv.Itoa(j+1)
			columnsSqlInsert[j] = columns[j].Name
			columnsSqlParams[j] = "$" + strconv.Itoa(j+1)
		}
		if !flagId {
			continue
		}
		gen.schema = append(gen.schema, TableInfo{
			Table:           tables[i],
			Columns:         columns,
			SqlColumnInsert: strings.Join(columnsSqlInsert, ", "),
			SqlColumnParams: strings.Join(columnsSqlParams, ", "),
			SqlColumnUpdate: strings.Join(columnsSqlUpdate, ", "),
		})
	}
	return nil
}

func (gen *Generator) saveFile() error {
	dir := "src/typ"
	if _, err := os.Stat(dir); err != nil {
		if err := os.Mkdir(dir, 0700); err != nil {
			return err
		}
	}
	const delimetr = "// BEFORE CODE GENERATED. DO NOT EDIT //"

	_, currentFile, _, _ := runtime.Caller(0)
	path := filepath.Dir(currentFile)

	for i := range gen.schema {
		variable := map[string]interface{}{
			"schema": gen.schema[i],
		}
		buf, err := tpl.ParseFile(path+"/typ.tpl", nil, variable)
		if err != nil {
			return err
		}
		//
		fileGo := dir + "/" + gen.schema[i].Table.Name + ".go"
		fi, err := os.Stat(fileGo)
		if err == nil && fi.Mode().IsRegular() {
			var data []byte
			if data, err = ioutil.ReadFile(fileGo); err != nil {
				return err
			}
			content := strings.Split(string(data), delimetr)
			buf.WriteString(content[1])
		}
		err = ioutil.WriteFile(fileGo, buf.Bytes(), 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

func (gen *Generator) typeConvert(col *Column) (string, string) {
	const NO = "NO"
	if elm, ok := gen.cfg.Generate.NameMapping[col.TableName+"."+col.Name]; ok {
		if col.IsNullable == NO {
			return elm.NotNull, elm.Tag
		} else {
			return elm.Null, elm.Tag
		}
	}
	if elm, ok := gen.cfg.Generate.NameMapping[col.Name]; ok {
		if col.IsNullable == NO {
			return elm.NotNull, elm.Tag
		} else {
			return elm.Null, elm.Tag
		}
	}
	if elm, ok := gen.cfg.Generate.TypeMapping[col.TypeFromDB]; ok {
		if col.IsNullable == NO {
			return elm.NotNull, elm.Tag
		} else {
			return elm.Null, elm.Tag
		}
	}
	if elm, ok := gen.TypeMappingDef[col.TypeFromDB][col.IsNullable]; ok {
		return elm, gen.TypeMappingDef[col.TypeFromDB]["Tag"]
	} else {
		return "null.JSON", ` swaggertype:"string"`
	}
}

type Config struct {
	Postgresql stpg.Config    `yaml:"postgresql"`
	Generate   GenerateConfig `json:"generate"`
}

type GenerateConfig struct {
	NameMapping map[string]PropertyConfig `yaml:"name_mapping"`
	TypeMapping map[string]PropertyConfig `yaml:"type_mapping"`
	TableIgnore []string                  `yaml:"table_ignore"`
	TableCRUD   []string                  `yaml:"table_crud"`
}

type PropertyConfig struct {
	Null    string `yaml:"null_yes"`
	NotNull string `yaml:"null_not"`
	Tag     string `yaml:"tag"`
}

type TableInfo struct {
	Table           *Table
	Columns         []*Column
	SqlColumnInsert string
	SqlColumnParams string
	SqlColumnUpdate string
}

type Table struct {
	Name        string      `db:"relname" json:"name"`
	NameType    string      `json:"name_type"`
	Description null.String `db:"description" json:"description"`
}

type Column struct {
	TableName   string      `db:"table_name" json:"table_name"`
	Name        string      `db:"column_name" json:"column_name"`
	NameType    string      `json:"name_type"`
	TypeFromDB  string      `db:"udt_name" json:"type_from_db"`
	TypeFromGO  string      `json:"type_from_go"`
	TagFromGO   string      `json:"tag_from_go"`
	IsNullable  string      `db:"is_nullable" json:"is_nullable"`
	Default     null.String `db:"column_default" json:"default"`
	Description null.String `db:"description" json:"description"`
}

const schema = "public"

const SQLTable = `
SELECT
	cl.relname,
	pd.description
FROM
	pg_catalog.pg_namespace AS n
INNER JOIN pg_catalog.pg_class AS cl ON
	n."oid" = cl.relnamespace
INNER JOIN information_schema."tables" t ON
	t.table_name = cl.relname
	AND t.table_schema = n.nspname
LEFT JOIN pg_catalog.pg_description pd ON
	cl."oid" = pd .objoid AND pd.objsubid = 0
WHERE
	n.nspname = $1
ORDER BY
	cl.relname ASC
`
const SQLColumn = `
SELECT
	col.table_name,
	col.column_name,
	col.udt_name,
	col.is_nullable,
	col.column_default,
	descp.description
FROM
	information_schema."columns" AS col
LEFT JOIN (
		SELECT
			pd.description,
			pd.objsubid
		FROM
			pg_catalog.pg_namespace AS n
		INNER JOIN pg_catalog.pg_class AS cl ON
			n."oid" = cl.relnamespace
		INNER JOIN pg_catalog.pg_description pd ON
			cl."oid" = pd .objoid
		WHERE
			n.nspname = $1
			AND cl.relname = $2
	) AS descp ON
	col.dtd_identifier::int = descp.objsubid::int
WHERE
	col.table_schema = $1
	AND col.table_name = $2
ORDER BY
	col.table_name,
	col.ordinal_position ASC
`
