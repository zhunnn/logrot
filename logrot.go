package logrot

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger struct {
	*log.Logger
	Setting setting
}

type setting struct {
	Format string
}

func NewLogger(name string, file bool, console bool, path string) (logger *Logger) {
	// Writer
	var writer io.Writer
	switch {
	case file && console:
		logFile, err := os.OpenFile(path+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		writer = io.MultiWriter(os.Stdout, logFile)
	case file:
		logFile, err := os.OpenFile(path+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		writer = io.MultiWriter(logFile)
	case console:
		writer = os.Stdout
	}
	// Logger
	name = fmt.Sprintf("[%v] ", name)
	logger = &Logger{Logger: log.New(writer, name, log.LstdFlags|log.Lshortfile)}

	return logger
}
