package logging

import (
	"blog/pkg/file"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

var (
	LevelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	logger     *log.Logger
	logPrefix  string
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

func Setup() {
	LogFileName := GetLogFileName()

	f, err := file.MustOpen(LogFileName, "runtime")
	if err != nil {
		log.Fatalf("open log file error: %v", err)
	}
	logger = log.New(f, "", log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}
func Warn(v ...interface{}) {
	setPrefix(WARN)
	logger.Println(v)
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, f, line, ok := runtime.Caller(2)

	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", LevelFlags[level], filepath.Base(f), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", LevelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
