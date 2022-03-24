package errs

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

var traceAllow string

func init() {
	d, _ := os.Getwd()
	if path.Base(d) == "bin" {
		d = path.Dir(d)
	}
	traceAllow = path.Dir(d)
}

func Traces() []string {
	tr := make([]string, 0, 10)
	for i := 4; true; i++ {
		t, _, _, _ := Trace(i)
		if t == "" {
			break
		}
		tr = append(tr, t)
	}

	return tr
}

func Trace(step int) (kind, filePath, funcName string, line int) {
	pc, filePath, line, ok := runtime.Caller(step)
	if line == 0 || (traceAllow != "" && !strings.Contains(filePath, traceAllow)) {
		return
	}

	filePath = strings.ReplaceAll(filePath, traceAllow, "")
	kind = fmt.Sprintf("%s:%d ", filePath, line)

	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			funcName = path.Base(fn.Name())
			kind += funcName
		}
	}
	return
}
