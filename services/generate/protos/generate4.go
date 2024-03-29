// Инструмент по автоматизации сопоставления типов для GRPC
//
// Генерация описаний прототипов, самих прототипов и методов конвертации типа в обе стороны.
// Обрабатываются только публичные и помеченные тегом json поля структур.
//
// Из коробки обрабатывает базовые типы golang (string, bool, int..., uint..., float..., []byte, []string)
// + uuid.UUID - реализация работы с полями UUID
// + time.Time - дата и время
// + time.Duration - время
// + decimal.Decimal - работа с дробными числами
// + ENUM в парадигме GRPC
// + ссылки на другие типы в этом же пакете (typs)
// + срезы ссылок на другие типы в этом же пакете (typs)
// + имеет спецификацию работы с типами используемыми в библиотеке boiler
// (null.JSON, null.Bytes, null.String, null.Time, types.StringArray)

package protos

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func Generate4(dir, md, pb string) {
	var err error
	var tplPFull, tplMFull, tplP, tplM string
	gen := Generate{
		controlType: map[string]bool{},
	}

	tplPFull = CreateProtoMessageFile(dir, pb)
	tplMFull = CreateMethodTypeFile(md)
	for _, t := range GenerateConfig[md] {
		if tplP, tplM, err = gen.ParseType(t, md, pb); err != nil {
			log.Fatal(err)
		}
		tplPFull += tplP
		tplMFull += tplM
	}
	// service proto (описание сервиса)
	if err = os.WriteFile(dir+"/"+pb+"/"+serviceName+".proto", []byte(tplPFull), 0o600); err != nil {
		log.Fatal(err)
	}
	// golang методы конвертации
	if err = os.WriteFile(dir+"/"+md+"/protom.go", []byte(tplMFull), 0o600); err != nil {
		log.Fatal(err)
	}
	// вспомогательные функции реализующие уникальную обработку свойств для определяемых рабочих типов
	d, err := os.ReadFile(dir + "/generate/data/protof.go")
	if err != nil {
		log.Fatal(err)
	}
	d = []byte(strings.ReplaceAll(string(d), "package data", "package "+md))
	if err := os.WriteFile(dir+"/"+md+"/protof.go", d, 0o600); err != nil {
		log.Fatal(err)
	}
	//
	_ = os.Remove(dir + "/generate/protos/config_work.go")
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
	tplMFrom := "\nfunc New" + list[1] + FromProto + "(proto *" + pkgProto + "." + list[1] + ") *" + list[1] + " {\n"
	tplMFrom += "\tif proto == nil { return nil }\n"
	tplMFrom += "\treturn &" + list[1] + "{\n"

	// one object type to proto
	tplMTo := "\nfunc New" + list[1] + ToProto + "(tt *" + list[1] + ") *" + pkgProto + "." + list[1] + " {\n"
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
	tplP += "\nmessage " + list[1] + SuffixSlice + " {\n\trepeated " + list[1] + " slice = 1;\n}\n"

	// slice proto to type
	tplMFrom += "\t}\n}\n\n" + gen.GenerateFuncSliceProtoType(list[1], pkgProto)

	// slice type to proto
	tplMTo += "\t}\n}\n\n" + gen.GenerateFuncSliceTypeProto(list[1], pkgProto)

	return tplP, tplMTo + tplMFrom, nil
}

// GenerateFuncSliceTypeProto генерация метода конвертации среза типа в срез его прототипа
func (gen *Generate) GenerateFuncSliceTypeProto(typ, pkgProto string) (s string) {
	s += fmt.Sprintf("func New%s%s"+ToProto+" (tt []*%s) []*%s.%s {", typ, SuffixSlice, typ, pkgProto, typ)
	s += fmt.Sprintf("\n\tres := make([]*%s.%s, len(tt))", pkgProto, typ)
	s += "\n\tfor i := range tt {"
	s += fmt.Sprintf("\n\t\tres[i] = New%s"+ToProto+"(tt[i])", typ)
	return s + "\n\t}\n\treturn res\n}\n"
}

// GenerateFuncSliceProtoType генерация метода конвертации среза прототипа в соответсвующий ему срез типа
func (gen *Generate) GenerateFuncSliceProtoType(typ, pkgProto string) (s string) {
	s += fmt.Sprintf("func New%s%s"+FromProto+"(list []*%s.%s) []*%s {", typ, SuffixSlice, pkgProto, typ, typ)
	s += fmt.Sprintf("\n\tres := make([]*%s, len(list))", typ)
	s += "\n\tfor i := range list {"
	s += fmt.Sprintf("\n\t\tres[i] = New%s"+FromProto+"(list[i])", typ)
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

	if f, ok := CustomHandlerFunc[propType]; ok {
		return f(i, fieldName, fieldJSON)
	}

	switch propKind {
	case reflect.String:
		if strings.Contains(propType, "enum.") {
			tplP, tplMFrom, tplMTo = GenerateFieldEnum(i, propType, fieldName, fieldJSON)
		} else {
			tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
			tplMFrom, tplMTo = GenerateFieldNative(fieldName, fieldJSON)
		}

	case reflect.Bool:
		tplP += "\tbool " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = GenerateFieldNative(fieldName, fieldJSON)

	case reflect.Float32:
		tplP += "\tfloat " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = GenerateFieldNative(fieldName, fieldJSON)
	case reflect.Float64:
		tplP += "\tdouble " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = GenerateFieldNative(fieldName, fieldJSON)

	case reflect.Int:
		tplP, tplMFrom, tplMTo = GenerateFieldInt(i, fieldName, fieldJSON)
	case reflect.Int8:
		tplP, tplMFrom, tplMTo = GenerateFieldInt8(i, fieldName, fieldJSON)
	case reflect.Int16:
		tplP, tplMFrom, tplMTo = GenerateFieldInt16(i, fieldName, fieldJSON)
	case reflect.Int32:
		tplP += "\tint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = GenerateFieldNative(fieldName, fieldJSON)
	case reflect.Int64:
		if propType == "time.Duration" {
			tplP += "\tint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
			tplMFrom, tplMTo = GenerateTimeDuration(fieldName, fieldJSON)
		} else {
			tplP += "\tint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
			tplMFrom, tplMTo = GenerateFieldNative(fieldName, fieldJSON)
		}
	case reflect.Uint:
		tplP, tplMFrom, tplMTo = GenerateFieldUint(i, fieldName, fieldJSON)
	case reflect.Uint8:
		tplP, tplMFrom, tplMTo = GenerateFieldUint8(i, fieldName, fieldJSON)
	case reflect.Uint16:
		tplP, tplMFrom, tplMTo = GenerateFieldUint16(i, fieldName, fieldJSON)
	case reflect.Uint32:
		tplP += "\tuint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = GenerateFieldNative(fieldName, fieldJSON)
	case reflect.Uint64:
		tplP += "\tuint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
		tplMFrom, tplMTo = GenerateFieldNative(fieldName, fieldJSON)

	case reflect.Slice:
		typParse := strings.Split(propType, pkgType+".")
		if propType == "[]string" {
			tplP, tplMFrom, tplMTo = GenerateFieldStringArray(i, fieldName, fieldJSON)
		} else if propType == "[]uint8" {
			tplP, tplMFrom, tplMTo = GenerateFieldBytes(i, fieldName, fieldJSON)
		} else if len(typParse) == 2 {
			typParseAdv := strings.Split(typParse[1], SuffixSlice)
			if len(typParseAdv) == 2 {
				tplP, tplMFrom, tplMTo = GenerateFieldSlicePtrType(i, typParseAdv[0], fieldName, fieldJSON)
			} else if typParse[0] == "[]*" {
				tplP, tplMFrom, tplMTo = GenerateFieldSlicePtrType(i, typParse[1], fieldName, fieldJSON)
			} else {
				fmt.Println(subjErr)
			}
		} else {
			fmt.Println(subjErr)
		}

	case reflect.Struct:
		typParse := strings.Split(propType, pkgType+".")
		if len(typParse) == 2 {
			tplP, tplMFrom, tplMTo = GenerateFieldStructType(i, typParse[1], fieldName, fieldJSON)
		} else {
			fmt.Println(subjErr)
		}

	case reflect.Ptr:
		typParse := strings.Split(propType, "*"+pkgType+".")
		if len(typParse) == 2 {
			tplP, tplMFrom, tplMTo = GenerateFieldPtrType(i, typParse[1], fieldName, fieldJSON)
		} else {
			fmt.Println(subjErr)
		}

	default:
		fmt.Println(subjErr)
	}

	return tplP, tplMFrom, tplMTo
}
