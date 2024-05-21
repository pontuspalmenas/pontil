package pontil

import (
	"github.com/joho/godotenv"
	"os"
)

func MustGetEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("environment not set: " + name)
	}
	return v
}

func LoadEnv(env string) {
	OrPanic(godotenv.Load(env + ".env"))
}
