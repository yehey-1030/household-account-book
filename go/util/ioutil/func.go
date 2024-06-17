package ioutil

import (
	"runtime"
	"strings"
)

// FuncName return function name of FuncName() caller
// in format: "{package}.{receiver}.{methodName}"
// ex) "app.(*application).Get"
func FuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc).Name()
	return f[strings.LastIndex(f, "/")+1:]
}
