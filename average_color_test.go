package average_color

import (
	"image"
	"image/color"
	"log"
	"os"
	"testing"
)

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func getImage(imgPath string) image.Image {
	f, err := os.Open(imgPath)
	check(err)
	defer f.Close()
	img, _, err := image.Decode(f)
	check(err)
	return img
}

func testImage(t *testing.T, imgPath string, checkColor color.NRGBA) {
	img := getImage(imgPath)
	averageColor := AverageColor(img)
	if averageColor != checkColor {
		t.Errorf("%s: expected %v, got %v", imgPath, checkColor, averageColor)
	}
}

func TestAverageColor(t *testing.T) {
	testImage(t, "./testdata/white.png", color.NRGBA{255, 255, 255, 255})
	testImage(t, "./testdata/black.png", color.NRGBA{0, 0, 0, 255})
	testImage(t, "./testdata/marshmallow-night.png", color.NRGBA{32, 37, 57, 255})
	testImage(t, "./testdata/transparent-white.png", color.NRGBA{0, 0, 0, 0})
	testImage(t, "./testdata/transparent-noise.png", color.NRGBA{0, 0, 0, 0})
	testImage(t, "./testdata/marshmallow-night-translucent.png", color.NRGBA{32, 37, 57, 127})
	testImage(t, "./testdata/brbw.png", color.NRGBA{128, 64, 128, 255})
	testImage(t, "./testdata/trbw.png", color.NRGBA{170, 85, 170, 191})
	testImage(t, "./testdata/brbt.png", color.NRGBA{85, 0, 85, 191})
	testImage(t, "./testdata/audio-headset.png", color.NRGBA{53, 52, 52, 62})
	testImage(t, "./testdata/thumb.jpg", color.NRGBA{126, 140, 148, 255})
}
