package helpers

import (
	"path/filepath"
	"runtime"

	"github.com/rs/zerolog/log"
)

func GetCallerFuncName(funcDepth int) string {

	if funcDepth > 32 {
		log.Warn().
			Str("source", "GetCallerFuncName").
			Any("data", map[string]int{"funcDepth": funcDepth}).
			Msg("Function depth overflowed. Maximum allowed depth is 32")
		return ""
	}

	pc, _, _, _ := runtime.Caller(funcDepth)

	fn := runtime.FuncForPC(pc)

	return filepath.Base(fn.Name())
}
