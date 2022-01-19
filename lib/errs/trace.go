package errs

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

var traceAllow string

func SetTraceAllow(key string) {
	traceAllow = key
}

func Traces() []string {
	tr := make([]string, 0, 10)
	for i := 4; true; i++ {
		t := trace(i)
		if t == "" {
			break
		}
		tr = append(tr, t)
	}

	return tr
}

func trace(step int) string {
	pc, file, line, ok := runtime.Caller(step)
	if line == 0 || (traceAllow != "" && !strings.Contains(file, traceAllow)) {
		return ""
	}

	kind := fmt.Sprintf("%s:%d ", file, line)

	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			kind += path.Base(fn.Name())
		}
	}

	return kind
}

func getMessage(args []interface{}) string {
	switch len(args) {
	case 0:
		return ""
	case 1:
		return args[0].(string)
	default:
		return fmt.Sprintf(args[0].(string), args[1:]...)
	}
}
