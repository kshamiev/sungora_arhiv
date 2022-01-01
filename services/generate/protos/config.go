package protos

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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
	SuffixSlice = "Slice"                                   // config prefix slice types
	FromProto   = "FromProto"                               // config suffix name func
	ToProto     = "ToProto"                                 // config suffix name func
	Separator   = "// AFTER CODE GENERATED. DO NOT EDIT //" // separator code generation
)

// GenerateConfig Список типов для которых нужно реализовать работу по GRPC. Наполняется программно.
// При первой работе удаляются все файлы кроме этого.
var GenerateConfig = map[string][]interface{}{}

var serviceName string

func Init() (stepRun int, dir, pkgType, pkgProto string) {
	step := flag.String("step", "sample-1", "generate step: serviceName-stepNumber(1-4)")
	flag.Parse()
	fmt.Println(*step)
	l := strings.Split(*step, "-")
	if len(l) != 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	n, err := strconv.Atoi(l[1])
	if err != nil || n < 1 || n > 4 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	serviceName = l[0]
	_, currentFile, _, _ := runtime.Caller(0)
	return n, filepath.Dir(filepath.Dir(filepath.Dir(currentFile))), "md" + serviceName, "pb" + serviceName
}
