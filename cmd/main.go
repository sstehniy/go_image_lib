package main

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/sstehniy/go_image_lib/pkg/img2ascii"
)

func main() {
	absPath, _ := filepath.Abs("./assets/images/14996914924_f9380a07df_c.jpg")
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

	img2ascii.ConvertImageToAscii(img, 0.1)
}
