// Инструмент по автоматизации сопоставления типов для GRPC
//
// Генерация описаний прототипов, самих прототипов и методов конвертации типа в обе стороны.
// Обрабатываются только публичные и помеченные тегом json поля структур.
//
// Из коробки обрабатывает базовые типы golang (string, bool, int..., uint..., float..., []byte, []string)
// + typ.UUID - реализация работы с полями UUID
// + time.Time - дата и время
// + time.Duration - время
// + decimal.Decimal - работа с дробными числами
// + ENUM в парадигме GRPC
// + ссылки на другие типы в этом же пакете (typs)
// + срезы ссылок на другие типы в этом же пакете (typs)
// + имеет спецификацию работы с типами используемыми в библиотеке boiler
// (null.JSON, null.Bytes, null.String, null.Time, types.StringArray)
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"

	_ "sungora/models/generate/config"
	"sungora/models/generate/protos"
)

func main() {
	dir, md, pb := protos.Init()

	var err error
	var tplSFull, tplPFull, tplMFull, tplP, tplM string
	gen := Generate{
		controlType: map[string]bool{},
	}

	tplSFull = protos.CreateProtoServiceFile(dir, pb)
	tplPFull = protos.CreateProtoMessageFile(dir, pb)
	tplMFull = protos.CreateMethodTypeFile(md)
	for _, t := range protos.GenerateConfig[md] {
		if tplP, tplM, err = gen.ParseType(t, md, pb); err != nil {
			log.Fatal(err)
		}
		tplPFull += tplP
		tplMFull += tplM
	}
	// type proto (описание прототипов)
	if err = ioutil.WriteFile(dir+"/"+pb+"/types.proto", []byte(tplPFull), 0600); err != nil {
		log.Fatal(err)
	}
	// service proto (описание сервиса)
	if err = ioutil.WriteFile(dir+"/"+pb+"/service.proto", []byte(tplSFull), 0600); err != nil {
		log.Fatal(err)
	}
	// golang методы конвертации
	if err = ioutil.WriteFile(dir+"/"+md+"/proto_method.go", []byte(tplMFull), 0600); err != nil {
		log.Fatal(err)
	}
	// вспомогательные функции реализующие уникальную обработку свойств для определяемых рабочих типов
	d, err := ioutil.ReadFile(dir + "/generate/protos/proto_func.go")
	if err != nil {
		log.Fatal(err)
	}
	d = []byte(strings.ReplaceAll(string(d), "package protos", "package "+md))
	if err := ioutil.WriteFile(dir+"/"+md+"/proto_func.go", d, 0600); err != nil {
		log.Fatal(err)
	}
}

// //// TYPE

type Generate struct {
	controlType map[string]bool
}

// ParseType Анализируем тип и формируем его сопряжение с grpc (proto файлы и методы конвертации) (Object = *TypeName)
func (gen *Generate) ParseType(object interface{}, pkgType, pkgProto string) (tplP, tplM string, err error) {
	// разбираем тип
	var value = reflect.ValueOf(object)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	list := strings.Split(value.Type().String(), ".")
	if _, ok := gen.controlType[list[1]]; ok {
		return tplP, tplM, nil
	}
	gen.controlType[list[1]] = true

	// pb
	tplP = "\nmessage " + list[1] + " {\n"

	// one object proto to type
	tplMFrom := "\nfunc New" + list[1] + protos.SuffixFromProto + "(proto *" + pkgProto + "." + list[1] + ") *" + list[1] + " {\n"
	tplMFrom += "\tif proto == nil { return nil }\n"
	tplMFrom += "\treturn &" + list[1] + "{\n"

	// one object type to proto
	tplMTo := "\nfunc New" + list[1] + protos.SuffixToProto + "(tt *" + list[1] + ") *" + pkgProto + "." + list[1] + " {\n"
	tplMTo += "\tif tt == nil { return nil }\n"
	tplMTo += "\treturn &" + pkgProto + "." + list[1] + "{\n"

	// разбираем свойства типа
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		// пропускаем приватные свойства
		if !field.IsValid() || !field.CanSet() {
			continue
		}
		tplP_, tplMFrom_, tplMTo_ := gen.ParseField(value, i, pkgType)
		tplP += tplP_
		tplMFrom += tplMFrom_
		tplMTo += tplMTo_
	}
	tplP += "}\n"
	tplP += "\nmessage " + list[1] + protos.SuffixSlice + " {\n\trepeated " + list[1] + " slice = 1;\n}\n"

	// slice proto to type
	tplMFrom += "\t}\n}\n\n" + gen.GenerateFuncSliceProtoType(list[1], pkgProto)

	// slice type to proto
	tplMTo += "\t}\n}\n\n" + gen.GenerateFuncSliceTypeProto(list[1], pkgProto)

	return tplP, tplMTo + tplMFrom, nil
}

// GenerateFuncSliceTypeProto генерация метода конвертации среза типа в срез его прототипа
func (gen *Generate) GenerateFuncSliceTypeProto(typ string, pkgProto string) (s string) {
	s += fmt.Sprintf("func New%s%s"+protos.SuffixToProto+" (tt []*%s) []*%s.%s {", typ, protos.SuffixSlice, typ, pkgProto, typ)
	s += fmt.Sprintf("\n\tres := make([]*%s.%s, len(tt))", pkgProto, typ)
	s += "\n\tfor i := range tt {"
	s += fmt.Sprintf("\n\t\tres[i] = New%s"+protos.SuffixToProto+"(tt[i])", typ)
	return s + "\n\t}\n\treturn res\n}\n"
}

// GenerateFuncSliceProtoType генерация метода конвертации среза прототипа в соответсвующий ему срез типа
func (gen *Generate) GenerateFuncSliceProtoType(typ string, pkgProto string) (s string) {
	s += fmt.Sprintf("func New%s%s"+protos.SuffixFromProto+"(list []*%s.%s) []*%s {", typ, protos.SuffixSlice, pkgProto, typ, typ)
	s += fmt.Sprintf("\n\tres := make([]*%s, len(list))", typ)
	s += "\n\tfor i := range list {"
	s += fmt.Sprintf("\n\t\tres[i] = New%s"+protos.SuffixFromProto+"(list[i])", typ)
	return s + "\n\t}\n\treturn res\n}\n"
}

// //// FIELD

// ParseField Анализируем свойство типа и генерируем описание его прототипа и методы конвертации в обе стороны
// nolint
func (gen *Generate) ParseField(objValue reflect.Value, i int, pkgType string) (tplP, tplMFrom, tplMTo string) {
	field := objValue.Type().Field(i)
	fieldName := field.Name
	fieldJSON := field.Tag.Get(`json`)

	// пропускаем исключенные и не обозначенные свойства
	if fieldJSON == `-` || fieldJSON == "" {
		return
	}
	fieldJSON = strings.Split(fieldJSON, ",")[0]

	// формируем согласно типу
	prop := objValue.Field(i)
	propType := prop.Type().String()
	propKind := prop.Type().Kind()
	subjErr := "not implemented undefined property: %s.%s [%s] %s"
	subjErr = fmt.Sprintf(subjErr, objValue.Type().String(), fieldName, propKind, propType)

	if f, ok := protos.CustomHandlerFunc[propType]; ok {
		return f(i, fieldName, fieldJSON)
	}

	switch propKind {
	case reflect.String:
		tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = protos.GenerateFieldNative(fieldName, fieldJSON)

	case reflect.Bool:
		tplP += "\tbool " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = protos.GenerateFieldNative(fieldName, fieldJSON)

	case reflect.Float32:
		tplP += "\tfloat " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = protos.GenerateFieldNative(fieldName, fieldJSON)
	case reflect.Float64:
		tplP += "\tdouble " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = protos.GenerateFieldNative(fieldName, fieldJSON)

	case reflect.Int:
		tplP, tplMFrom, tplMTo = protos.GenerateFieldInt(i, fieldName, fieldJSON)
	case reflect.Int8:
		tplP, tplMFrom, tplMTo = protos.GenerateFieldInt8(i, fieldName, fieldJSON)
	case reflect.Int16:
		tplP, tplMFrom, tplMTo = protos.GenerateFieldInt16(i, fieldName, fieldJSON)
	case reflect.Int32:
		tplP += "\tint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = protos.GenerateFieldNative(fieldName, fieldJSON)
	case reflect.Int64:
		if propType == "time.Duration" {
			tplP += "\tint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
			tplMFrom, tplMTo = protos.GenerateTimeDuration(fieldName, fieldJSON)
		} else {
			tplP += "\tint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
			tplMFrom, tplMTo = protos.GenerateFieldNative(fieldName, fieldJSON)
		}
	case reflect.Uint:
		tplP, tplMFrom, tplMTo = protos.GenerateFieldUint(i, fieldName, fieldJSON)
	case reflect.Uint8:
		tplP, tplMFrom, tplMTo = protos.GenerateFieldUint8(i, fieldName, fieldJSON)
	case reflect.Uint16:
		tplP, tplMFrom, tplMTo = protos.GenerateFieldUint16(i, fieldName, fieldJSON)
	case reflect.Uint32:
		tplP += "\tuint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = protos.GenerateFieldNative(fieldName, fieldJSON)
	case reflect.Uint64:
		tplP += "\tuint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = protos.GenerateFieldNative(fieldName, fieldJSON)

	case reflect.Slice:
		typParse := strings.Split(propType, pkgType+".")
		if propType == "[]string" {
			tplP, tplMFrom, tplMTo = protos.GenerateFieldStringArray(i, fieldName, fieldJSON)
		} else if propType == "[]uint8" {
			tplP, tplMFrom, tplMTo = protos.GenerateFieldBytes(i, fieldName, fieldJSON)
		} else if len(typParse) == 2 {
			typParseAdv := strings.Split(typParse[1], protos.SuffixSlice)
			if len(typParseAdv) == 2 {
				tplP, tplMFrom, tplMTo = protos.GenerateFieldSlicePtrType(i, typParseAdv[0], fieldName, fieldJSON)
			} else if typParse[0] == "[]*" {
				tplP, tplMFrom, tplMTo = protos.GenerateFieldSlicePtrType(i, typParse[1], fieldName, fieldJSON)
			} else {
				fmt.Println(subjErr)
			}
		} else {
			fmt.Println(subjErr)
		}

	case reflect.Struct:
		typParse := strings.Split(propType, pkgType+".")
		if len(typParse) == 2 {
			tplP, tplMFrom, tplMTo = protos.GenerateFieldStructType(i, typParse[1], fieldName, fieldJSON)
		} else {
			fmt.Println(subjErr)
		}

	case reflect.Ptr:
		typParse := strings.Split(propType, "*"+pkgType+".")
		if len(typParse) == 2 {
			tplP, tplMFrom, tplMTo = protos.GenerateFieldPtrType(i, typParse[1], fieldName, fieldJSON)
		} else {
			fmt.Println(subjErr)
		}

	default:
		fmt.Println(subjErr)
	}

	return tplP, tplMFrom, tplMTo
}
