package main

import (
	"fmt"
	"github.com/Dercraker/MapTilesCutter/internal/cutter"
	"github.com/Dercraker/MapTilesCutter/internal/validator"
	"github.com/spf13/cobra"
	"log"
)

var (
	picturePath string
	outDir      string
	concurrency int

	cliCmd = &cobra.Command{
		Use:   "mapTilesCutter",
		Short: "A fast map tiles cutter in go",
		Long:  "This tools can cut File in many tiles for Map Tiled System like googleMap or Leaflet",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runMapCutter(picturePath, outDir); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	cliCmd.Flags().StringVarP(&picturePath, "picturePath", "f", "./map.png", "The path of input picture")
	cliCmd.Flags().StringVarP(&outDir, "outputPath", "o", "./tiles/", "The path of output tiles generated from picture")
	cliCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "The number of concurrent goroutines for tile processing")
}

func main() {
	if err := cliCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func runMapCutter(file, outDir string) error {
	validateArgs(file, outDir)
	if err := cutter.ProcessFile(&file, &outDir, &concurrency); err != nil {
		return err
	}
	return nil
}

func validateArgs(file, outDir string) {
	fmt.Println("==========================")
	fmt.Println("Validating File")
	fmt.Println("==========================")
	validator.IsFileExist(file)
	fmt.Println("IsFileExist : Done")
	validator.FileIsReadable(file)
	fmt.Println("FileIsReadable : Done")
	validator.FileIsCorrectSize(file)
	fmt.Println("FileIsCorrectSize : Done")
	validator.FileHasCorrectType(file)
	fmt.Println("FileHasCorrectType : Done")

	fmt.Println("==========================")
	fmt.Println("Validating Folder")
	fmt.Println("==========================")
	validator.IsFolderExist(outDir)
	fmt.Println("IsFolderExist : Done")
	validator.IsFolderWritable(outDir)
	fmt.Println("IsFolderWritable : Done")
}
