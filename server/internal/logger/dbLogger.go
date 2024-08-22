package logger

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	gormLogger "gorm.io/gorm/logger"
)

type DBLogger struct {
	log zerolog.Logger
	gormLogger.Config
}

type requestIdType string

const RequestIdKey = requestIdType("requestId")

func NewDBLogger(log zerolog.Logger, config gormLogger.Config) gormLogger.Interface {
	return &DBLogger{
		log:    log,
		Config: config,
	}
}

func (l *DBLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l *DBLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Info {
		l.log.Info().Msgf(msg, data...)
	}
}

func (l *DBLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Warn {
		l.log.Warn().Msgf(msg, data...)
	}
}

func (l *DBLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Error {
		l.log.Error().Msgf(msg, data...)
	}
}

func (l *DBLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	requestId, ok := ctx.Value(RequestIdKey).(string)

	if !ok {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	event := l.log.With().
		Str("requestId", requestId).
		Dur("queryTime", elapsed).
		Str("source", "dbLogger").
		Str("sql", sql).
		Int64("rows", rows).
		Logger()

	event.Trace().Send()
}
