package img2ascii

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
)

const (
	gscale1 = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~i!lI;:,\"^`"
	gscale2 = "@%#*+=-:. "
)

func ConvertImageToAscii(img image.Image, scale float64) {
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	width := img.Bounds().Max.X - img.Bounds().Min.X
	colorsArray := make([][]color.Color, height)
	for i := 0; i < len(colorsArray); i++ {
		colorsArray[i] = make([]color.Color, width)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			colorsArray[i][j] = img.At(j, i)
		}
	}
	fmt.Println("1")
	// print old height and width
	fmt.Println(height, width)

	newWidth := int(float64(width) * scale)
	newHeight := int(float64(height) * scale)
	// print new height and width
	fmt.Println(newHeight, newWidth)
	brightnessMatrix := make([][]byte, newHeight)
	for i := 0; i < len(brightnessMatrix); i++ {
		brightnessMatrix[i] = make([]byte, newWidth)
	}

	fmt.Println("2")
	for i := 0; i < newHeight; i++ {
		for j := 0; j < newWidth; j++ {
			x1 := int(float64(i) / scale)
			y1 := int(float64(j) / scale)
			x2 := int(float64(i+1) / scale)
			y2 := int(float64(j+1) / scale)
			brightnessMatrix[i][j] = byte(getAvarageL(&img, x1, y1, x2, y2) / 256 / float64(len(gscale1)))
		}
	}
	fmt.Println("3")
	fmt.Println(len(brightnessMatrix), len(brightnessMatrix[0]))
	for i := 0; i < newHeight; i++ {
		for j := 0; j < newWidth; j++ {
			fmt.Print(string(gscale1[brightnessMatrix[i][j]]))
		}
		fmt.Print("\n")
	}

}

func getAvarageL(img *image.Image, x1 int, y1 int, x2 int, y2 int) float64 {
	var sum float64 = 0
	maxX := (*img).Bounds().Max.X
	maxY := (*img).Bounds().Max.Y
	if x2 > maxX {
		x2 = maxX
	}
	if y2 > maxY {
		y2 = maxY
	}
	for i := x1; i < x2; i++ {
		for j := y1; j < y2; j++ {
			r, g, b, _ := (*img).At(j, i).RGBA()
			sum += 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
		}
	}
	sum /= float64((x2 - x1) * (y2 - y1))
	return sum
}
