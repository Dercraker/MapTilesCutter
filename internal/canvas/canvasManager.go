package canvas

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
)

func CreateCanvas(zoom, tileSize int, color color.Color) image.Image {
	width := tileSize * (1 << zoom)
	canvas := image.NewRGBA(image.Rect(0, 0, width, width))

	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			canvas.Set(x, y, color)
		}
	}

	return canvas
}

func MergeImgToCanvas(canvas, img image.Image) image.Image {
	canvasBounds := canvas.Bounds()
	imgBounds := img.Bounds()

	x := (canvasBounds.Dx() - imgBounds.Dx()) / 2
	y := (canvasBounds.Dy() - imgBounds.Dy()) / 2

	return imaging.Paste(canvas, img, image.Pt(x, y))
}
