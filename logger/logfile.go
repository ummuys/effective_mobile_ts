package logger

import (
	"fmt"
	"os"
	"time"
)

type currentLog struct {
	date string
	file *os.File
}

func initLogFile(path string) *currentLog {
	date := time.Now().Format("2006-01-02")
	os.MkdirAll(path, 0755)

	file, err := os.OpenFile(fmt.Sprintf("%s/%s.log", path, date), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Errorf("can't create/open file: %w", err))
	}
	return &currentLog{
		date: date,
		file: file,
	}
}

func (l *currentLog) Write(p []byte) (n int, err error) {
	date := time.Now().Format("2006-01-02")
	if date != l.date {
		l.file.Close()
		file, err := os.OpenFile(fmt.Sprintf("logs/%s.log", date), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Errorf("can't create/open file: %w", err))
		}
		l.date = date
		l.file = file
	}
	return l.file.Write(p)
}
