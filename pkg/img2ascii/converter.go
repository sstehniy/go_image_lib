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

func NewAsciiConverter(img image.Image, detailed bool, scale float64) *AsciiConverter {
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	width := img.Bounds().Max.X - img.Bounds().Min.X
	numPixels := height * width
	// Calculate scale based on a logarithmic function for smoother transitions
	if scale == 0.0 {

		scale = math.Log10(float64(numPixels)) / (math.Log10(float64(numPixels)) + 15) // Adjust the denominator based on desired scaling
	}

	grayscaleArray := image.NewGray(img.Bounds())
	draw.Draw(grayscaleArray, grayscaleArray.Bounds(), img, img.Bounds().Min, draw.Src)
	avgContrast := calcAvgContrast(grayscaleArray)
	println("avgContrast:", avgContrast)

	return &AsciiConverter{
		Image:    img,
		Scale:    scale,
		Width:    width,
		Height:   height,
		Detailed: detailed,
		GScaled:  grayscaleArray,
	}
}

func (c AsciiConverter) GetWidth() int {
	return c.Width
}

func (c AsciiConverter) GetHeight() int {
	return c.Height
}

func (c AsciiConverter) GetScale() float64 {
	return c.Scale
}

func (c AsciiConverter) GetScaledDims() (int, int) {
	return int(float64(c.Width) * c.Scale), int(float64(c.Height) * c.Scale)
}

func (c *AsciiConverter) WithScale(scale float64) *AsciiConverter {
	c.Scale = scale
	return c
}

func (c *AsciiConverter) WithDetailed(detailed bool) *AsciiConverter {
	c.Detailed = detailed
	return c
}

func (c *AsciiConverter) Convert() string {
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
	outputString := ""
	gscale := gscale1
	if c.Detailed == false {
		gscale = gscale2
	}
	for i := 0; i < scaledHeight; i++ {
		for j := 0; j < scaledWidth; j++ {
			idx := int(float64(grayScaled.GrayAt(j, i).Y) / 255 * float64(len(gscale)-1))
			outputString += string(gscale[idx])
		}
		outputString += "\n"
	}

	return outputString
}
