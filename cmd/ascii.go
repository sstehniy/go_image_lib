package cmd

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sstehniy/gopix/pkg/img2ascii"
)

var (
	input    string
	output   string
	scale    float64
	detailed bool
)

var asciiCmd = &cobra.Command{
	Use:   "ascii",
	Short: "gopix is a CLI tool for converting images to ASCII art... and more!",
	Run: func(cmd *cobra.Command, args []string) {

		absPath, _ := filepath.Abs(input)
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
		println("Detailed:", detailed)
		println("Scale:", scale)
		converter := img2ascii.NewAsciiConverter(img, detailed, scale)
		outputStr := converter.Convert()
		if len(output) == 0 {
			println()
			println(outputStr)
		}

		extension := filepath.Ext(output)
		fmt.Println(strings.ToLower(extension))

		if strings.ToLower(extension) != ".txt" && strings.ToLower(extension) != ".png" {
			fmt.Println("Please choose a valid file extension (.txt or .png)")
			os.Exit(1)
		}
		file, err = os.Create(output)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		switch extension {
		case ".txt":
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			file.WriteString(outputStr)
		case ".png":
			imgBytes, err := img2ascii.ConvertToPNG(outputStr, converter.GetWidth(), converter.GetHeight())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			file.Write(imgBytes)
		}

	},
}

func init() {
	asciiCmd.Flags().StringVarP(&input, "input", "i", "", "Input file path")
	asciiCmd.Flags().StringVarP(&output, "output", "o", "", "Output file name")
	asciiCmd.Flags().Float64VarP(&scale, "scale", "s", 0.0, "Scale of the ASCII image")
	asciiCmd.Flags().BoolVarP(&detailed, "detailed", "d", false, "Use detailed ASCII characters")
}
