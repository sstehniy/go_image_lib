package img2ascii

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	fontSize = 12
)

func ConvertToPNG(text string, width, height int) ([]byte, error) {
	fgColor := color.RGBA{255, 255, 255, 255}
	bgColor := color.RGBA{0, 0, 0, 255}
	fg := image.NewUniform(fgColor)
	bg := image.NewUniform(bgColor)
	width = width * fontSize
	height = height * fontSize
	fontPath, err := filepath.Abs("./assets/fonts/Montserrat-Regular.ttf")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// print width, height as struct

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	ctx := freetype.NewContext()
	draw.Draw(img, img.Bounds(), bg, image.Pt(0, 0), draw.Src)

	fontParsed, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	ctx.SetFont(fontParsed)
	ctx.SetDPI(72)
	ctx.SetFontSize(fontSize)
	ctx.SetClip(img.Bounds())
	ctx.SetDst(img)
	ctx.SetSrc(fg)
	ctx.SetHinting(font.HintingNone)
	textXOffset := 0

	textSplitted := strings.Split(text, "\n")
	chunkSize := 1000
	wg := &sync.WaitGroup{}
	for i := 0; i < len(textSplitted); i += chunkSize {
		wg.Add(1)
		end := i + chunkSize
		if end > len(textSplitted) {
			end = len(textSplitted)
		}
		chunk := textSplitted[i:end]

		go func(chunk []string, idx int) {
			defer wg.Done()

			yOffset := int(ctx.PointToFixed(float64(idx+1)*fontSize*2.22) >> 6)
			pt := freetype.Pt(textXOffset, yOffset)
			for _, s := range chunk {
				initialX := pt.X
				for _, c := range s {
					_, err := ctx.DrawString(strings.Replace(string(c), "\r", "", -1), pt)
					if err != nil {
						return
					}
					pt.X += ctx.PointToFixed(fontSize)
				}
				pt.X = initialX
				pt.Y += ctx.PointToFixed(fontSize * 2.22)
			}
		}(chunk, i/chunkSize)
	}

	wg.Wait()

	b := new(bytes.Buffer)
	if err := png.Encode(b, img); err != nil {
		log.Println("unable to encode image.")
		return nil, err
	}
	return b.Bytes(), nil
}
