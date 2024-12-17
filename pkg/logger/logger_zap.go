package logger

import (
	"go.uber.org/zap"
	"log"
)

type Logger struct {
	log *zap.Logger
}

func MustInitZap(deb bool) *Logger {
	var err error
	var logger *zap.Logger
	if deb {
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Fatalf("Ошибка инициализации логгера разработки")
		}
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			log.Fatalf("Ошибка инициализации логгера")
		}
	}
	return &Logger{log: logger}
}

func (z *Logger) Info(msg string, fields ...any) {
	z.log.Sugar().Infow(msg, fields...)
}

func (z *Logger) Error(msg string, fields ...any) {
	z.log.Sugar().Errorw(msg, fields...)
}

func (z *Logger) Debug(msg string, fields ...any) {
	z.log.Sugar().Debugw(msg, fields...)
}
