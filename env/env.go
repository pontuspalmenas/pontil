package env

import (
	"github.com/joho/godotenv"
	"github.com/pontuspalmenas/pontil"
	"os"
)

func MustGet(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("environment not set: " + name)
	}
	return v
}

func GetOrDefault(name string, def string) string {
	v := os.Getenv(name)
	if v == "" {
		return def
	}
	return v
}

func Load(env string) {
	pontil.OrPanic(godotenv.Load(env + ".env"))
}
