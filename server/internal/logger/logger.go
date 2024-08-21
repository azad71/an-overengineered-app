package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var once sync.Once

var log zerolog.Logger

func GetLogger() *zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logLevel = int(zerolog.TraceLevel)
		}

		var outputChannel io.Writer = prettyJSONWriter{os.Stdout}

		if os.Getenv("APP_ENV") != "local" {
			outputChannel = os.Stderr
		}

		log = zerolog.New(outputChannel).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Logger()

		zerolog.DefaultContextLogger = &log
	})

	return &log
}

func SetContext(log *zerolog.Logger, data map[string]string) {
	for key, value := range data {
		log.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str(key, value)
		})
	}
}

type prettyJSONWriter struct {
	out io.Writer
}

func (w prettyJSONWriter) Write(p []byte) (n int, err error) {
	var jsonObj map[string]interface{}
	if err := json.Unmarshal(p, &jsonObj); err != nil {
		return w.out.Write(p) // Write raw data if unmarshaling fails
	}

	// Determine the color based on the log level
	level := jsonObj["level"].(string)
	color := getColorForLevel(level)

	// Pretty-print JSON
	prettyJSON, err := json.MarshalIndent(jsonObj, "", "    ")
	if err != nil {
		return w.out.Write(p) // Write raw data if marshaling fails
	}

	// Convert the pretty-printed JSON to a string and apply the color
	coloredStr := fmt.Sprintf("%s%s%s\n", color, string(prettyJSON), "\033[0m")

	// Write the colorized JSON string to the output
	return w.out.Write([]byte(coloredStr))
}

// getColorForLevel returns the ANSI color code based on the log level
func getColorForLevel(level string) string {
	switch level {
	case "info":
		return "\033[1;32m" // Green
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
