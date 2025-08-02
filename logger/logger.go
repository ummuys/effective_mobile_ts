package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

func InitLogger(path string) *zerolog.Logger {

	//STD-OUT
	file := initLogFile(path)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}

	multiWriter := io.MultiWriter(file, consoleWriter)

	logger := zerolog.New(multiWriter).With().Timestamp().Logger()
	return &logger
}
