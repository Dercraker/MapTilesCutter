package validator

import (
	"log"
	"os"
	"path/filepath"
)

func IsFolderExist(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		log.Println("Output dir does not exist")

		if err := os.MkdirAll(dir, 0750); err != nil {
			log.Fatalln("Fail to create Output dir")
		}
		log.Println("Output dir created")
	}
}

func IsFolderWritable(dir string) bool {
	testFile := filepath.Join(dir, ".tmp_write_test")
	f, err := os.Create(testFile)
	if err != nil {
		return false
	}
	f.Close()
	os.Remove(testFile)
	return true
}
