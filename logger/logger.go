package logger

import "github.com/rs/zerolog"

func InitLogger(path string) *zerolog.Logger {
	logger := zerolog.New(initLogFile(path)).With().Timestamp().Logger()
	return &logger
}
