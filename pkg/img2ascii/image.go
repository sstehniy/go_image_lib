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
	fontSize = 100
)

func ConvertToPNG(text string) ([]byte, error) {
	fgColor := color.RGBA{0, 255, 255, 255}
	bgColor := color.RGBA{255, 0, 0, 255}
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

	width := len(strings.Split(text, "\n")[0])
	height := len(strings.Split(text, "\n"))
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
	fmt.Println(ctx.PointToFixed(fontSize))
	ctx.SetClip(img.Bounds())
	ctx.SetDst(img)
	ctx.SetSrc(fg)
	ctx.SetHinting(font.HintingNone)
	textXOffset := 50
	textYOffset := 10 + int(ctx.PointToFixed(fontSize)>>6)

	pt := freetype.Pt(textXOffset, textYOffset)
	for _, s := range text {
		_, err := ctx.DrawString(strings.Replace(string(s), "\r", "", -1), pt)
		if err != nil {
			return nil, err
		}
		pt.Y += ctx.PointToFixed(fontSize * 1.5)
	}

	b := new(bytes.Buffer)
	if err := png.Encode(b, img); err != nil {
		log.Println("unable to encode image.")
		return nil, err
	}
	return b.Bytes(), nil
}
