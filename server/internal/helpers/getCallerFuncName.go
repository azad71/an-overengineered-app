package helpers

import (
	"path/filepath"
	"runtime"
)

func GetCallerFuncName(funcDepth int) string {
	pc, _, _, _ := runtime.Caller(funcDepth)

	fn := runtime.FuncForPC(pc)

	return filepath.Base(fn.Name())
}
