package zlog

import (
	"cmp"
	"fmt"
	"os"
)

func Trace(s string) {
	if cmp.Or(os.Getenv("LOG_LEVEL"), "default") == "trace" {
		writeLog("[TRACE] " + s + "\n")
	}
}

func Debug(s string) {
	writeLog("[DEBUG] " + s + "\n")
}

func Debugf(format string, a ...any) {
	Debug(fmt.Sprintf(format, a...))
}

func Info(s string) {
	writeLog("[INFO] " + s + "\n")
}

func Infof(format string, a ...any) {
	Info(fmt.Sprintf(format, a...))
}

func Error(s string) {
	writeLog("[ERROR] " + s + "\n")
}

func Errorf(format string, a ...any) {
	Error(fmt.Sprintf(format, a...))
}

func writeLog(s string) {
	fmt.Print(s)
}
