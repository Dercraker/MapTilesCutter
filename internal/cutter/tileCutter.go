package cutter

import (
	"fmt"
	canvasManager "github.com/Dercraker/MapTilesCutter/internal/canvas"
	"github.com/Dercraker/MapTilesCutter/internal/handler"
	"github.com/disintegration/imaging"
	"github.com/schollz/progressbar/v3"
	"image"
	"image/color"
	"strconv"
	"sync"
)

type TileHandler interface {
	HandleTile(img image.Image, z, x, y int) error
}

type TileCutter struct {
	TileSize        int
	MinZoom         int
	MaxZoom         int
	BackgroundColor color.Color
	Concurrency     int
	Handler         TileHandler
	ProgessBar      *progressbar.ProgressBar
}

func ProcessFile(file, outDir *string, concurency *int) error {
	fmt.Println("Start Preprocessing configuration")

	fmt.Println("Opening File :", *file)
	srcFile, err := imaging.Open(*file)
	if err != nil {
		return err
	}
	fmt.Println("File Loaded")

	tileHandler := handler.NewFileHandler(*outDir)

	minZoom, maxZoom, totalTiles := getZoomLevels(srcFile, 256)

	tc := TileCutter{
		TileSize:        256,
		MinZoom:         minZoom,
		MaxZoom:         maxZoom,
		Concurrency:     *concurency,
		BackgroundColor: color.Black,
		Handler:         tileHandler,
		ProgessBar:      progressbar.Default(int64(totalTiles)),
	}

	fmt.Println("All preprocessing configuration done")
	fmt.Println("===================================================")
	fmt.Println("Minimum zoom level:", minZoom)
	fmt.Println("Maximum zoom level:", maxZoom)
	fmt.Println("Number of map tile to generate:", strconv.Itoa(totalTiles))
	fmt.Println("===================================================")

	if err := tc.CutMap(srcFile); err != nil {
		return err
	}

	fmt.Println("File has been cutted")
	return nil
}

func (tc *TileCutter) CutMap(img image.Image) error {
	var wg sync.WaitGroup
	sem := make(chan struct{}, tc.Concurrency)

	for z := tc.MinZoom; z <= tc.MaxZoom; z++ {
		canvas := canvasManager.CreateCanvas(z, tc.TileSize, tc.BackgroundColor)
		scaled := scaleImageToFit(img, canvas.Bounds())
		merged := canvasManager.MergeImgToCanvas(canvas, scaled)

		tilesX := merged.Bounds().Dx() / tc.TileSize
		tilesY := merged.Bounds().Dy() / tc.TileSize

		for x := 0; x < tilesX; x++ {
			for y := 0; y < tilesY; y++ {
				wg.Add(1)
				sem <- struct{}{}

				go func(z, x, y int, img image.Image) {
					defer func() {
						wg.Done()
						<-sem
					}()

					tile := cropTile(img, x, y, tc.TileSize)

					if err := tc.Handler.HandleTile(tile, z, x, y); err != nil {
						fmt.Printf("Tile processing error (%d, %d, %d): %v\n", z, x, y, err)
					}
					tc.ProgessBar.Add(1)
				}(z, x, y, merged)
			}
		}
	}
	wg.Wait()
	return nil
}
