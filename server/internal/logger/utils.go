package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"strings"
)

func getCallerFuncSource(funcDepth int) (string, string, error) {

	if funcDepth > 32 {
		return "", "", errors.New("maximum function depth exceeds")
	}

	ptr, file, line, _ := runtime.Caller(funcDepth)

	funcName := strings.Split(filepath.Base(runtime.FuncForPC(ptr).Name()), ".")[1]

	return fmt.Sprintf("%s:%d", file, line), filepath.Base(funcName), nil
}

// ANSI color code based on the log level
func getColorForLevel(level string) string {
	switch level {
	case "info":
		return "\033[0;32m" // Green
	case "error":
		return "\033[1;31m" // Red
	case "fatal":
		return "\033[1;36m" // Cyan
	case "panic":
		return "\033[1;34m" // blue
	default:
		return "\033[0m" // Reset/No color
	}
}

type PrettyJSONWriter struct {
	Out io.Writer
}

func (w PrettyJSONWriter) Write(p []byte) (n int, err error) {
	var jsonObj map[string]interface{}
	originalLen := len(p) // This is to satisfy Write func contract and avoid "short write" error

	if err := json.Unmarshal(p, &jsonObj); err != nil {
		// Write raw data if unmarshaling fails
		w.Out.Write(p)
		return originalLen, err
	}

	level, _ := jsonObj["level"].(string)
	color := getColorForLevel(level)

	prettyJSON, err := json.MarshalIndent(jsonObj, "", "  ")
	if err != nil {

		w.Out.Write(p)
		return originalLen, err
	}

	coloredStr := fmt.Sprintf("%s%s%s\n", color, string(prettyJSON), "\033[0m")

	w.Out.Write([]byte(coloredStr))

	// Return original length, even if output differs
	// It doesn't matter though, the objective is to push the log to standard output channel
	return originalLen, err
}
