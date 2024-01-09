package clog

import (
	"runtime"
	"strings"
)

func GetCaller(callDepth int, suffixesToIgnore ...string) (file string, line int, pc uintptr) {
	// bump by 1 to ignore the getCaller (this) stackframe
	callDepth++
outer:
	for {
		var ok bool
		pc, file, line, ok = runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
			break
		}

		for _, s := range suffixesToIgnore {
			if strings.HasSuffix(file, s) {
				callDepth++
				continue outer
			}
		}
		break
	}
	return
}

// GetCallerIgnoringLogMulti TODO
func GetCallerIgnoringLogMulti(callDepth int) (string, int, uintptr) {
	// the +1 is to ignore this (getCallerIgnoringLogMulti) frame
	return GetCaller(callDepth+1, "/hooks.go",
		"/entry.go", "/logger.go", "/exported.go", "clog/clog.go", "asm_amd64.s")
}

func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}
