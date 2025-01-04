package logger

import (
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var (
	once sync.Once
	log  zerolog.Logger
)

func GetLogger() *zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano
		zerolog.MessageFieldName = "msg"
		zerolog.TimestampFieldName = "ts"

		logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logLevel = int(zerolog.TraceLevel)
		}

		var outputChannel io.Writer = os.Stdout
		if os.Getenv("APP_ENV") == "local" {
			outputChannel = PrettyJSONWriter{os.Stdout}
		}

		log = zerolog.
			New(outputChannel).
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
