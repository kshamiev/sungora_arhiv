package protos

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

// generation functions matching proto type and real type
// функции реализующие обработку пользовательских типов свойств для определяемых рабочих типов
// описание в протофайлах, приведение значений в обе стороны, вызов вспомогательных функций для обработким

var CustomHandlerFunc = map[string]func(int, string, string) (string, string, string){
	"decimal.Decimal":   GenerateFieldDecimal,
	"[]typ.UUID":        GenerateFieldUUIDSlice,
	"typ.UUIDS":         GenerateFieldUUIDSlice,
	"types.StringArray": GenerateFieldStringArray,
	"typ.UUID":          GenerateFieldUUID,
	"time.Time":         GenerateFieldTime,
	"null.Time":         GenerateFieldNullTime,
	"null.String":       GenerateFieldNullString,
	"null.Bytes":        GenerateFieldNullBytes,
	"null.JSON":         GenerateFieldNullJSON,
}

// GenerateFieldNullJSON конвертация - сопоставление туда и обратно
func GenerateFieldNullJSON(i int, pType, pMessage string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tbytes " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("\t\t%s: tt.%s.JSON,\n", ConvFP(pMessage), pType)
	tplMFrom = fmt.Sprintf("\t\t%s: pbFromNullJSON(proto.%s),\n", pType, ConvFP(pMessage))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldNullBytes конвертация - сопоставление туда и обратно
func GenerateFieldNullBytes(i int, pType, pMessage string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tbytes " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tt.%s.Bytes,\n", ConvFP(pMessage), pType)
	tplMFrom = fmt.Sprintf("%s: pbFromNullBytes(proto.%s),\n", pType, ConvFP(pMessage))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldNullString конвертация - сопоставление туда и обратно
func GenerateFieldNullString(i int, pType, pMessage string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tstring " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tt.%s.String,\n", ConvFP(pMessage), pType)
	tplMFrom = fmt.Sprintf("%s: pbFromNullString(proto.%s),\n", pType, ConvFP(pMessage))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldNullTime конвертация - сопоставление туда и обратно
func GenerateFieldNullTime(i int, pType, pMessage string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tgoogle.protobuf.Timestamp " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: pbToTime(tt.%s.Time),\n", ConvFP(pMessage), pType)
	tplMFrom = fmt.Sprintf("%s: pbFromNullTime(proto.%s),\n", pType, ConvFP(pMessage))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldTime конвертация - сопоставление туда и обратно
func GenerateFieldTime(i int, pType, pMessage string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tgoogle.protobuf.Timestamp " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: pbToTime(tt.%s),\n", ConvFP(pMessage), pType)
	tplMFrom = fmt.Sprintf("%s: pbFromTime(proto.%s),\n", pType, ConvFP(pMessage))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldUUID конвертация - сопоставление туда и обратно
func GenerateFieldUUID(i int, pType, pMessage string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tstring " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tt.%s.String(),\n", ConvFP(pMessage), pType)
	tplMFrom = fmt.Sprintf("%s: typ.UUIDMustParse(proto.%s),\n", pType, ConvFP(pMessage))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldDecimal конвертация - сопоставление туда и обратно
func GenerateFieldDecimal(i int, pType, pMessage string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tstring " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tt.%s.String(),\n", ConvFP(pMessage), pType)
	tplMFrom = fmt.Sprintf("%s: pbFromDecimal(proto.%s),\n", pType, ConvFP(pMessage))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldUUIDSlice конвертация - сопоставление туда и обратно
func GenerateFieldUUIDSlice(i int, pType, pMessage string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\trepeated string " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: pbToUUIDS(tt.%s),\n", ConvFP(pMessage), pType)
	tplMFrom = fmt.Sprintf("%s: pbFromUUIDS(proto.%s),\n", pType, ConvFP(pMessage))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldStringArray конвертация - сопоставление туда и обратно
func GenerateFieldStringArray(i int, pType, pMessage string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\trepeated string " + pMessage + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tt.%s,\n", ConvFP(pMessage), pType)
	tplMFrom = fmt.Sprintf("%s: proto.%s,\n", pType, ConvFP(pMessage))
	return tplP, tplMFrom, tplMTo
}

// //// стандартные и системные типы

// GenerateFieldNative конвертация - сопоставление туда и обратно
func GenerateFieldNative(field, fieldJSON string) (tplMFrom, tplMTo string) {
	tplMTo = fmt.Sprintf("%s: tt.%s,\n", ConvFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: proto.%s,\n", field, ConvFP(fieldJSON))
	return tplMFrom, tplMTo
}

// GenerateFieldPtrType конвертация - сопоставление туда и обратно
func GenerateFieldPtrType(i int, typParse, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\t" + typParse + " " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: New%s"+ToProto+"(tt.%s),\n", ConvFP(fieldJSON), typParse, field)
	tplMFrom = fmt.Sprintf("%s: New%s"+FromProto+"(proto.%s),\n", field, typParse, ConvFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldStructType конвертация - сопоставление туда и обратно
func GenerateFieldStructType(i int, typParse, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\t" + typParse + " " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: New%s"+ToProto+"(&tt.%s),\n", ConvFP(fieldJSON), typParse, field)
	tplMFrom = fmt.Sprintf("%s: *New%s"+FromProto+"(proto.%s),\n", field, typParse, ConvFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldSlicePtrType конвертация - сопоставление туда и обратно
func GenerateFieldSlicePtrType(i int, typParse, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\trepeated " + typParse + " " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: New%s%s"+ToProto+"(tt.%s),\n", ConvFP(fieldJSON), typParse, SuffixSlice, field)
	tplMFrom = fmt.Sprintf("%s: New%s%s"+FromProto+"(proto.%s),\n", field, typParse, SuffixSlice, ConvFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateTimeDuration конвертация - сопоставление туда и обратно
func GenerateTimeDuration(field, fieldJSON string) (tplMFrom, tplMTo string) {
	tplMTo = fmt.Sprintf("%s: tt.%s.Nanoseconds(),\n", ConvFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: time.Duration(proto.%s),\n", field, ConvFP(fieldJSON))
	return tplMFrom, tplMTo
}

// GenerateFieldUint8 конвертация - сопоставление туда и обратно
func GenerateFieldUint8(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tuint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: uint32(tt.%s),\n", ConvFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: uint8(proto.%s),\n", field, ConvFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldUint16 конвертация - сопоставление туда и обратно
func GenerateFieldUint16(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tuint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: uint32(tt.%s),\n", ConvFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: uint16(proto.%s),\n", field, ConvFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldUint конвертация - сопоставление туда и обратно
func GenerateFieldUint(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tuint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: uint64(tt.%s),\n", ConvFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: uint(proto.%s),\n", field, ConvFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldInt8 конвертация - сопоставление туда и обратно
func GenerateFieldInt8(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: int32(tt.%s),\n", ConvFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: int8(proto.%s),\n", field, ConvFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldInt16 конвертация - сопоставление туда и обратно
func GenerateFieldInt16(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tint32 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: int32(tt.%s),\n", ConvFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: int16(proto.%s),\n", field, ConvFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldInt конвертация - сопоставление туда и обратно
func GenerateFieldInt(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tint64 " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: int64(tt.%s),\n", ConvFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: int(proto.%s),\n", field, ConvFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldBytes конвертация - сопоставление туда и обратно
func GenerateFieldBytes(i int, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tbytes " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tt.%s,\n", ConvFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: proto.%s,\n", field, ConvFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// GenerateFieldEnum конвертация - сопоставление туда и обратно
func GenerateFieldEnum(i int, typParse, field, fieldJSON string) (tplP, tplMFrom, tplMTo string) {
	tplP += "\tstring " + fieldJSON + " = " + strconv.Itoa(i+1) + ";\n"
	tplMTo = fmt.Sprintf("%s: tt.%s.String(),\n", ConvFP(fieldJSON), field)
	tplMFrom = fmt.Sprintf("%s: %s(proto.%s),\n", field, typParse, ConvFP(fieldJSON))
	return tplP, tplMFrom, tplMTo
}

// //// служебные методы

// ConvFP Получение названия свойства в прототипе (через тег json) для сопоставления
func ConvFP(fieldTag string) string {
	if fieldTag == "id" {
		return "Id"
	}
	list := strings.Split(fieldTag, "_")
	for i := range list {
		list[i] = strings.Title(list[i])
	}
	return strings.Join(list, "")
}

// CreateProtoMessageFile инициализация файла с описанием прототипов
func CreateProtoMessageFile(d, pkgProto string) (proto string) {
	if data, err := ioutil.ReadFile(d + "/" + pkgProto + "/types.proto"); err == nil {
		proto = string(data)
		list := strings.Split(proto, Separator)
		proto = list[0] + Separator + "\n"
	} else {
		if data, err = ioutil.ReadFile(d + "/generate/data/types.proto"); err != nil {
			log.Fatal(err)
		}
		proto = string(data)
		proto = strings.ReplaceAll(proto, "TPLpackage", pkgProto)
	}
	return proto
}

// CreateProtoServiceFile инициализация файла с описанием сервиса
func CreateProtoServiceFile(d, pkgProto string) (proto string) {
	if data, err := ioutil.ReadFile(d + "/" + pkgProto + "/service.proto"); err == nil {
		proto = string(data)
	} else {
		if data, err = ioutil.ReadFile(d + "/generate/data/service.proto"); err != nil {
			log.Fatal(err)
		}
		proto = string(data)
		proto = strings.ReplaceAll(proto, "TPLpackage", pkgProto)
		proto = strings.ReplaceAll(proto, "TPLservice", strings.Title(strings.Split(pkgProto, "-")[1]))
		proto = strings.ReplaceAll(proto, "types.proto", filepath.Base(d)+"/"+pkgProto+"/types.proto")
	}
	return proto
}

// CreateMethodTypeFile инициализация файла с методами конвертации (protobuf)
func CreateMethodTypeFile(pkgType string) (proto string) {
	return "package " + pkgType + "\n\n"
}
