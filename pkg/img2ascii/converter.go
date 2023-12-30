package img2ascii

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"math"
)

const (
	gscale1                = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~i!lI;:,\"^`"
	gscale2                = "@%#*+=-:. "
	IMAGE_HEIGHT_REDUCTION = 0.45
)

type AsciiConverter struct {
	Image    image.Image
	Scale    float64
	Detailed bool
	Width    int
	Height   int
	GScaled  *image.Gray
}

func NewAsciiConverter(img image.Image) *AsciiConverter {
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	width := img.Bounds().Max.X - img.Bounds().Min.X
	numPixels := height * width
	var scale float64 = 1
	if numPixels > 20000 {
		scale = math.Log10(float64(numPixels)) / 10
	}
	if numPixels > 100000 {
		scale = math.Log10(float64(numPixels)) / 25
	}
	if numPixels > 200000 {
		scale = math.Log10(float64(numPixels)) / 35
	}
	if numPixels > 500000 {
		scale = math.Log10(float64(numPixels)) / 50
	}
	if numPixels > 1000000 {
		scale = math.Log10(float64(numPixels)) / 75
	}
	if numPixels > 2000000 {
		scale = math.Log10(float64(numPixels)) / 100
	}
	if numPixels > 5000000 {
		scale = math.Log10(float64(numPixels)) / 175
	}
	grayscaleArray := image.NewGray(img.Bounds())
	draw.Draw(grayscaleArray, grayscaleArray.Bounds(), img, img.Bounds().Min, draw.Src)
	avgContrast := calcAvgContrast(grayscaleArray)
	println("avgContrast:", avgContrast)
	detailed := true
	if avgContrast < 90 {
		detailed = false
	}

	return &AsciiConverter{
		Image:    img,
		Scale:    scale,
		Width:    width,
		Height:   height,
		Detailed: detailed,
		GScaled:  grayscaleArray,
	}
}

func (c *AsciiConverter) WithScale(scale float64) *AsciiConverter {
	c.Scale = scale
	return c
}

func (c *AsciiConverter) WithDetailed(detailed bool) *AsciiConverter {
	c.Detailed = detailed
	return c
}

func (c *AsciiConverter) Convert() {
	if c.Image == nil {
		panic("Image must be set")
	}
	reducedHeight := int(float64(c.Height) * IMAGE_HEIGHT_REDUCTION)
	scaledHeight := int(float64(reducedHeight) * c.Scale)
	scaledWidth := int(float64(c.Width) * c.Scale)

	println("height:", c.Height, "width:", c.Width)
	println("reducedHeight:", reducedHeight, "scaledHeight:", scaledHeight, "scaledWidth:", scaledWidth)

	transformedColorsArray := make([][]color.Color, scaledHeight)
	for i := 0; i < len(transformedColorsArray); i++ {
		transformedColorsArray[i] = make([]color.Color, scaledWidth)
	}

	for i := 0; i < scaledHeight; i++ {
		for j := 0; j < scaledWidth; j++ {
			origY := int(float64(i) / IMAGE_HEIGHT_REDUCTION / c.Scale)
			origX := int(float64(j) / c.Scale)

			if origY >= c.Height {
				origY = c.Height - 1
			}
			if origX >= c.Width {
				origX = c.Width - 1
			}

			sampledPixel := c.GScaled.At(origX, origY)

			transformedColorsArray[i][j] = sampledPixel
		}
	}

	transformedImage := convertMatrixToImage(transformedColorsArray)

	grayScaled := image.NewGray(image.Rect(0, 0, scaledWidth, scaledHeight))
	draw.Draw(grayScaled, grayScaled.Bounds(), transformedImage, transformedImage.Bounds().Min, draw.Src)
	gscale := gscale1
	if c.Detailed == false {
		gscale = gscale2
	}
	for i := 0; i < scaledHeight; i++ {
		for j := 0; j < scaledWidth; j++ {
			idx := int(float64(grayScaled.GrayAt(j, i).Y) / 255 * float64(len(gscale)-1))

			print(string(gscale[idx]))
		}
		print("\n")
	}
}

// func ConvertImageToAscii(img image.Image, scale float64) {
// 	height := img.Bounds().Max.Y - img.Bounds().Min.Y
// 	width := img.Bounds().Max.X - img.Bounds().Min.X

// 	grayscaleArray := image.NewGray(img.Bounds())
// 	draw.Draw(grayscaleArray, grayscaleArray.Bounds(), img, img.Bounds().Min, draw.Src)
// 	avgContrast := calcAvgContrast(grayscaleArray)
// 	println("avgContrast:", avgContrast)
// 	gscale := gscale1
// 	if avgContrast < 90 {
// 		gscale = gscale2
// 	}
// 	reducedHeight := int(float64(height) * IMAGE_HEIGHT_REDUCTION)
// 	scaledHeight := int(float64(reducedHeight) * scale)
// 	scaledWidth := int(float64(width) * scale)

// 	println("height:", height, "width:", width)
// 	println("reducedHeight:", reducedHeight, "scaledHeight:", scaledHeight, "scaledWidth:", scaledWidth)

// 	transformedColorsArray := make([][]color.Color, scaledHeight)
// 	for i := 0; i < len(transformedColorsArray); i++ {
// 		transformedColorsArray[i] = make([]color.Color, scaledWidth)
// 	}

// 	for i := 0; i < scaledHeight; i++ {
// 		for j := 0; j < scaledWidth; j++ {
// 			origY := int(float64(i) / IMAGE_HEIGHT_REDUCTION / scale)
// 			origX := int(float64(j) / scale)

// 			if origY >= height {
// 				origY = height - 1
// 			}
// 			if origX >= width {
// 				origX = width - 1
// 			}

// 			sampledPixel := grayscaleArray.At(origX, origY)

// 			transformedColorsArray[i][j] = sampledPixel
// 		}
// 	}

// 	transformedImage := convertMatrixToImage(transformedColorsArray)

// 	grayScaled := image.NewGray(image.Rect(0, 0, scaledWidth, scaledHeight))
// 	draw.Draw(grayScaled, grayScaled.Bounds(), transformedImage, transformedImage.Bounds().Min, draw.Src)

// 	for i := 0; i < scaledHeight; i++ {
// 		for j := 0; j < scaledWidth; j++ {
// 			idx := int(float64(grayScaled.GrayAt(j, i).Y) / 255 * float64(len(gscale)-1))

// 			print(string(gscale[idx]))
// 		}
// 		print("\n")
// 	}

// }
