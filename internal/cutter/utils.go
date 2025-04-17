package cutter

import (
	"image"

	"github.com/disintegration/imaging"
)

func scaleImageToFit(img image.Image, maxBounds image.Rectangle) image.Image {
	imgBounds := img.Bounds()
	imgWidth := imgBounds.Dx()
	imgHeight := imgBounds.Dy()

	ratioWidth := float64(maxBounds.Dx()) / float64(imgWidth)
	ratioHeight := float64(maxBounds.Dy()) / float64(imgHeight)

	ratio := ratioWidth
	if ratioHeight < ratioWidth {
		ratio = ratioHeight
	}

	newWidth := int(float64(imgWidth) * ratio)
	newHeight := int(float64(imgHeight) * ratio)

	return imaging.Resize(img, newWidth, newHeight, imaging.Linear)
}

func cropTile(img image.Image, x, y, tileSize int) image.Image {
	rect := image.Rect(x*tileSize, y*tileSize, (x+1)*tileSize, (y+1)*tileSize)
	return imaging.Crop(img, rect)
}
