package img2ascii

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

const (
	fontSize = 16
)

func ConvertToPNG(text string, width, height int) ([]byte, error) {
	fgColor := color.RGBA{255, 255, 255, 255}
	bgColor := color.RGBA{0, 0, 0, 255}
	width = width + 20
	fg := image.NewUniform(fgColor)
	bg := image.NewUniform(bgColor)
	fontPath, err := filepath.Abs("./assets/fonts/Montserrat-Regular.ttf")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fontBytes, err := os.ReadFile(fontPath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// print width, height as struct

	img := image.NewRGBA(image.Rect(0, 0, width*10, height*15))
	ctx := freetype.NewContext()
	draw.Draw(img, img.Bounds(), bg, image.Pt(0, 0), draw.Src)

	fontParsed, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	ctx.SetFont(fontParsed)
	ctx.SetDPI(72)
	ctx.SetFontSize(fontSize)
	fmt.Println(ctx.PointToFixed(fontSize))
	ctx.SetClip(img.Bounds())
	ctx.SetDst(img)
	ctx.SetSrc(fg)
	ctx.SetHinting(font.HintingNone)
	textXOffset := 0
	textYOffset := int(ctx.PointToFixed(fontSize) >> 6)

	pt := freetype.Pt(textXOffset, textYOffset)
	textSplitted := strings.Split(text, "\n")
	for _, s := range textSplitted {
		initialX := pt.X
		for _, c := range s {
			_, err := ctx.DrawString(strings.Replace(string(c), "\r", "", -1), pt)
			if err != nil {
				return nil, err
			}
			pt.X += ctx.PointToFixed(fontSize)
		}
		pt.X = initialX
		pt.Y += ctx.PointToFixed(fontSize * 1.5)
	}

	b := new(bytes.Buffer)
	if err := png.Encode(b, img); err != nil {
		log.Println("unable to encode image.")
		return nil, err
	}
	return b.Bytes(), nil
}
