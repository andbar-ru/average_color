package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/andbar-ru/average_color"
)

const (
	hexDesc = "output color in hex format, e.g. #21293e or #aa55aabf"
	rgbDesc = "output color in rgba format, e.g. rgb(33,41,62) or rgba(170,85,170,191)"
)

var (
	// Command line arguments
	hex, rgb  bool
	imagePath string
)

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func printHelpAndExit(code int) {
	fmt.Printf("Usage: %s [flags] <image path>\n", os.Args[0])
	fmt.Println("Flags:")
	fmt.Printf("  -hex (default) - %s\n", hexDesc)
	fmt.Printf("  -rgb - %s\n", rgbDesc)

	os.Exit(code)
}

func parseArgs() {
	if len(os.Args) == 1 {
		printHelpAndExit(0)
	}

	flag.BoolVar(&hex, "hex", false, hexDesc)
	flag.BoolVar(&rgb, "rgb", false, rgbDesc)

	flag.Parse()

	if !hex && !rgb {
		hex = true
	}

	imagePath = flag.Arg(0)
	if imagePath == "" {
		fmt.Println("ERROR! Image path is not given.\n")
		printHelpAndExit(1)
	}
}

func main() {
	parseArgs()

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		fmt.Printf("file %s does not exist\n", imagePath)
		os.Exit(1)
	}

	f, err := os.Open(imagePath)
	check(err)
	defer f.Close()
	img, _, err := image.Decode(f)
	check(err)

	averageColor := average_color.AverageColor(img)

	fmt.Println(imagePath)
	if averageColor.A == 0xff {
		if hex {
			fmt.Printf("#%02x%02x%02x", averageColor.R, averageColor.G, averageColor.B)
		}
		if rgb {
			if hex {
				fmt.Print(" ")
			}
			fmt.Printf("rgb(%d,%d,%d)", averageColor.R, averageColor.G, averageColor.B)
		}
	} else {
		if hex {
			fmt.Printf("#%02x%02x%02x%02x", averageColor.R, averageColor.G, averageColor.B, averageColor.A)
		}
		if rgb {
			if hex {
				fmt.Print(" ")
			}
			fmt.Printf("rgba(%d,%d,%d,%d)", averageColor.R, averageColor.G, averageColor.B, averageColor.A)
		}
	}
	fmt.Println()
}
