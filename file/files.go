package util

import (
	"os"
	"path/filepath"
)

func ForAllFilesInDir(path string, f func(file *os.File) error) error {
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			err = f(file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
