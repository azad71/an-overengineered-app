package logger

import (
	"context"

	"github.com/rs/zerolog"
)

func setLogMetaData(log *zerolog.Event) *zerolog.Event {
	src, funcName, _ := getCallerFuncSource(3)
	return log.Str("src", src).Str("func", funcName)
}

func Info(ctx context.Context, msg string, data any) {
	log := zerolog.Ctx(ctx).Info()
	log = setLogMetaData(log)

	if data != nil {
		log = log.Interface("data", data)
	}

	log.Msg(msg)
}

func Error(ctx context.Context, msg string, data any) {
	log := zerolog.Ctx(ctx).Error()
	log = setLogMetaData(log)

	log.Msg(msg)

}

func ErrorStack(ctx context.Context, msg string, err error, errData ...any) {
	if err == nil {
		return
	}

	log := zerolog.Ctx(ctx).Error()
	log = setLogMetaData(log).Err(err).Stack()

	if errData != nil {
		log.Interface("errData", errData)
	}

	if msg != "" {
		log.Msg(msg)
	} else {
		log.Send()
	}

}

func Fatal(ctx context.Context, msg string, err error) {
	log := zerolog.Ctx(ctx).Fatal()
	log = setLogMetaData(log)

	log.Err(err).Stack().Msg(msg)
}

func Panic(ctx context.Context, msg string, err error) {
	log := zerolog.Ctx(ctx).Panic()
	log = setLogMetaData(log)

	log.Err(err).Stack().Msg(msg)
}

func Warning(ctx context.Context, msg string, data any) {
	log := zerolog.Ctx(ctx).Warn()
	log = setLogMetaData(log)

	if data != nil {
		log.Interface("data", data)
	}

	log.Msg(msg)
}
