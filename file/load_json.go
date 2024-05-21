package util

import (
	"encoding/json"
	. "github.com/pontuspalmenas/pontil"
	"os"
	"path/filepath"
	"sync"
)

func LoadAllJsonConcurrent[T any](path string) []T {
	var objects []T
	var wg sync.WaitGroup
	var mutex sync.Mutex

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			wg.Add(1)
			go func(p string) {
				defer wg.Done()
				obj := loadJson[T](path)
				mutex.Lock()
				objects = append(objects, obj)
				mutex.Unlock()
			}(path)
		}
		return nil
	})
	OrPanic(err)
	wg.Wait()

	return objects
}

func loadJson[T any](path string) T {
	bs, err := os.ReadFile(path)
	OrPanic(err)

	var data T
	OrPanic(json.Unmarshal(bs, &data))
	return data
}
