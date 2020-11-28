package app

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

func Traces(err error) (tr []string) {
	kind := ""
	if err != nil {
		kind = err.Error() + "; "
	}

	for i := 2; true; i++ {
		t := Trace(i)
		if t == "" {
			break
		}

		switch {
		case strings.Contains(t, "/src/"): // LIBRARY GOPATH
			continue
		case strings.Contains(t, "/mod/"): // LIBRARY MOD
			continue
		case strings.Contains(t, "/vendor/"): // LIBRARY VENDOR
			continue
		}
		tr = append(tr, kind+t)
	}
	return tr
}

func Trace(step int) (kind string) {
	pc, file, line, ok := runtime.Caller(step)
	if line == 0 {
		return
	}
	kind = fmt.Sprintf("%s:%d", file, line)
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			kind += ":" + path.Base(fn.Name())
		}
	}
	return
}

func TraceAtom(step int) (kind, file, fun string, line int) {
	var ok bool
	var pc uintptr
	pc, file, line, ok = runtime.Caller(step)
	if line == 0 {
		return
	}
	kind = fmt.Sprintf("%s:%d", file, line)
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			fun = path.Base(fn.Name())
			kind += ":" + fun
		}
	}
	return
}
