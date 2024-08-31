package logger

import (
	"an-overengineered-app/internal/helpers"
	"context"

	"github.com/rs/zerolog"
)

func getLoggerInstance(ctx context.Context) *zerolog.Logger {

	if ctx == nil || ctx == context.TODO() {
		return GetLogger()
	} else {
		return zerolog.Ctx(ctx)
	}
}

func PrintInfo(ctx context.Context, msg string, data any) {

	if msg == "" {
		return
	}
	log := getLoggerInstance(ctx)

	callerFunc := helpers.GetCallerFuncName(2)

	infoLog := log.Info().Str("source", callerFunc)

	if data != nil {
		infoLog.Interface("data", data).Msg(msg)
	} else {
		infoLog.Msg(msg)
	}
}

func PrintError(ctx context.Context, msg string) {

	if msg == "" {
		return
	}

	callerFunc := helpers.GetCallerFuncName(2)

	log := getLoggerInstance(ctx)

	log.Error().Str("source", callerFunc).Msg(msg)

}

func PrintErrorWithStack(ctx context.Context, msg string, err error) {

	callerFunc := helpers.GetCallerFuncName(2)

	log := getLoggerInstance(ctx)

	if msg != "" {
		log.Error().Err(err).Stack().Str("source", callerFunc).Msg(msg)
	} else {
		log.Error().Err(err).Stack().Str("source", callerFunc).Send()
	}

}

func PrintFatal(ctx context.Context, msg string, err error) {
	callerFunc := helpers.GetCallerFuncName(2)

	log := getLoggerInstance(ctx)

	log.Fatal().Err(err).Stack().Str("source", callerFunc).Msg(msg)
}

func PrintPanic(ctx context.Context, msg string, err error) {
	callerFunc := helpers.GetCallerFuncName(2)

	log := getLoggerInstance(ctx)

	log.Panic().Err(err).Stack().Str("source", callerFunc).Msg(msg)
}

func PrintWarning(ctx context.Context, msg string, data any) {
	callerFunc := helpers.GetCallerFuncName(2)
	log := getLoggerInstance(ctx).Warn().Str("source", callerFunc)

	if data != nil {
		log.Interface("data", data)
	} else {
		log.Msg(msg)
	}

}
