package cutter

import (
	"image"
	"math"
)

func getZoomLevels(img image.Image, tileSize int) (minZoom, maxZoom, totalTiles int) {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	maxSize := math.Max(float64(width), float64(height))
	tilesAtMaxZoom := maxSize / float64(tileSize)

	maxZoom, totalTiles = 0, 0

	for math.Pow(2, float64(maxZoom)) < tilesAtMaxZoom {
		maxZoom++
	}

	for z := 0; z <= maxZoom; z++ {
		n := int(math.Pow(2, float64(z)))
		totalTiles += n * n
	}

	return
}
