package handler

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"os"
	"path/filepath"
)

type FileHandler struct {
	OutputDir string
}

func NewFileHandler(outputDir string) *FileHandler {
	return &FileHandler{
		OutputDir: outputDir,
	}
}

func (f *FileHandler) HandleTile(img image.Image, z, x, y int) error {
	dir := filepath.Join(f.OutputDir, fmt.Sprintf("%d", z), fmt.Sprintf("%d", x))
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	filename := filepath.Join(dir, fmt.Sprintf("%d.png", y))
	return imaging.Save(img, filename)
}
