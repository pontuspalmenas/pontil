package util

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

func Info(s string) {
	writeLog("[INFO] " + s + "\n")
}

func Error(s string) {
	writeLog("[ERROR] " + s + "\n")
}

func writeLog(s string) {
	fmt.Print(s)
}
