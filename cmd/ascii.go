package cmd

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sstehniy/gopix/pkg/img2ascii"
)

var asciiCmd = &cobra.Command{
	Use:   "ascii",
	Short: "gopix is a CLI tool for converting images to ASCII art... and more!",
	Run: func(cmd *cobra.Command, args []string) {
		paths := []string{
			"./assets/images/14996914924_f9380a07df_q.jpg",
			"./assets/images/p3286591407-5-800x533.jpg",
			"./assets/images/Silhouette-contrast-photos.png",
			"./assets/images/sunset-anime-comet-stars-scenery-digital-art-4k-wallpaper-uhdpaper.com-771@0@i.jpg",
		}

		for _, path := range paths {
			absPath, _ := filepath.Abs(path)
			file, err := os.Open(absPath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer file.Close()
			img, imageType, err := image.Decode(file)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Image type:", imageType)

			converter := img2ascii.NewAsciiConverter(img)
			converter.Convert()
		}
	},
}
