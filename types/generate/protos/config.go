package protos

import (
	"flag"
	"os"
	"path/filepath"
	"runtime"
)

// FileListCustom Файлы кодовую часть которых нужно сохранить в неизменном виде
// Прячем на время работы sqlboiler
var FileListCustom = []string{
	"hooks.go",
	"proto.go",
}

// FileException Файлы которые исключены из обработки
var FileException = map[string]bool{
	"boil_queries.go":     true,
	"boil_table_names.go": true,
	"boil_types.go":       true,
	"psql_upsert.go":      true,
	"hooks.go":            true,
}

const (
	SuffixSlice     = "Slice"                                   // config prefix slice types
	SuffixFromProto = "FromProto"                               // config suffix name func
	SuffixToProto   = "ToProto"                                 // config suffix name func
	Separator       = "// AFTER CODE GENERATED. DO NOT EDIT //" // separator code generation
)

// GenerateConfig Список типов для которых нужно реализовать работу по GRPC. Наполняется программно.
// При первой работе удаляются все файлы кроме этого.
var GenerateConfig = map[string][]interface{}{}

func Init() (string, string, string) {
	md := flag.String("md", "", "package type name (folder)")
	pb := flag.String("pb", "", "package proto name (folder)")
	flag.Parse()
	if *md == "" || *pb == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	_, currentFile, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filepath.Dir(currentFile))), *md, *pb
}
