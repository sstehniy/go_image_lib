package img2ascii

import (
	"image"
	"image/color"
	"testing"
)

func TestConvertMatrixToImage(t *testing.T) {
	// Define a test matrix
	matrix := [][]color.Color{
		{color.RGBA{255, 0, 0, 255}, color.RGBA{0, 255, 0, 255}},
		{color.RGBA{0, 0, 255, 255}, color.RGBA{255, 255, 255, 255}},
	}

	// Call the function
	img := convertMatrixToImage(matrix)

	// Check the dimensions of the image
	if img.Bounds().Dx() != 2 || img.Bounds().Dy() != 2 {
		t.Errorf("Expected image dimensions to be 2x2, but got %dx%d", img.Bounds().Dx(), img.Bounds().Dy())
	}

	// Check the colors of the image
	for i, row := range matrix {
		for j, col := range row {
			if img.At(j, i) != col {
				t.Errorf("Expected pixel at (%d, %d) to be %v, but got %v", j, i, col, img.At(j, i))
			}
		}
	}
}

func TestCalcAvgContrast(t *testing.T) {
	tests := []struct {
		name     string
		img      *image.Gray
		expected float64
	}{
		{
			name: "Test with uniform image",
			img: &image.Gray{
				Pix:    []uint8{128, 128, 128, 128, 128, 128, 128, 128, 128},
				Stride: 3,
				Rect:   image.Rect(0, 0, 3, 3),
			},
			expected: 0,
		},
		{
			name: "Test with varying image",
			img: &image.Gray{
				Pix:    []uint8{255, 0, 255, 0, 255, 0, 255, 0, 255},
				Stride: 3,
				Rect:   image.Rect(0, 0, 3, 3),
			},
			expected: 126.71051872498808,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcAvgContrast(tt.img); got != tt.expected {
				t.Errorf("calcAvgContrast() = %v, want %v", got, tt.expected)
			}
		})
	}
}
