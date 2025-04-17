package validator

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func IsFileExist(file string) {
	if _, err := os.Stat(file); err != nil {
		log.Fatalln(fmt.Sprintf("file %s not exist", file))
	}
}

func FileIsReadable(file string) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func FileIsCorrectSize(file string) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	img, _, err := image.DecodeConfig(f)
	if err != nil {
		log.Fatalln(err)
	}

	if (img.Width < 256) || (img.Height < 256) {
		log.Fatalln("Image is too small. The picture min size is 256*256px")
	}
}

func FileHasCorrectType(file string) {
	ext := strings.ToLower(filepath.Ext(file))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		log.Fatalln("File type is not supported")
	}
}
