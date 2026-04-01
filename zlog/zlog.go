package zlog

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	levelDebug = iota
	levelInfo
	levelWarn
	levelError
)

func logLevel() int {
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		return levelDebug
	case "info":
		return levelInfo
	case "warn":
		return levelWarn
	case "error":
		return levelError
	default:
		return levelInfo
	}
}

func Debug(msg string) {
	if logLevel() <= levelDebug {
		writeLog("[debug] " + msg + "\n")
	}
}

func Debugf(format string, a ...any) {
	if logLevel() <= levelDebug {
		writeLog("[debug] " + fmt.Sprintf(format, a...) + "\n")
	}
}

func Info(msg string) {
	if logLevel() <= levelInfo {
		writeLog("[info] " + msg + "\n")
	}
}

func Infof(format string, a ...any) {
	if logLevel() <= levelInfo {
		writeLog("[info] " + fmt.Sprintf(format, a...) + "\n")
	}
}

func Warn(msg string) {
	if logLevel() <= levelWarn {
		writeLog("[warn] " + msg + "\n")
	}
}

func Warnf(format string, a ...any) {
	if logLevel() <= levelWarn {
		writeLog("[warn] " + fmt.Sprintf(format, a...) + "\n")
	}
}

func Error(msg string) {
	if logLevel() <= levelError {
		writeLog("[error] " + msg + "\n")
	}
}

func Errorf(format string, a ...any) {
	if logLevel() <= levelError {
		writeLog("[error] " + fmt.Sprintf(format, a...) + "\n")
	}
}

func writeLog(s string) {
	fmt.Printf("%s %s", time.Now().Format(time.RFC3339), s)
}
