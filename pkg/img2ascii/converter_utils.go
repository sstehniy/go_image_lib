package img2ascii

import (
	"image"
	"image/color"
	"math"
)

func convertMatrixToImage(matrix [][]color.Color) image.Image {
	height, width := len(matrix), len(matrix[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			img.Set(j, i, matrix[i][j])
		}
	}
	return img
}

func calcAvgContrast(img *image.Gray) float64 {
	avgLum := 0.0
	for i := 0; i < len(img.Pix); i++ {
		avgLum += float64(img.Pix[i])
	}
	avgLum /= float64(len(img.Pix))
	varianceSum := 0.0
	for i := 0; i < len(img.Pix); i++ {
		diff := float64(img.Pix[i]) - avgLum
		varianceSum += diff * diff
	}
	variance := varianceSum / float64(len(img.Pix))

	return math.Sqrt(variance)
}
