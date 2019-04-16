package average_color

import (
	"image"
	"image/color"
	"math"
	"sync"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type colorSums struct {
	mux sync.Mutex
	// sum of alpha-premultiplied color values ∈ [0,65535]
	redSum   uint64
	greenSum uint64
	blueSum  uint64
	alphaSum uint64
}

func (cs *colorSums) inc(r, g, b, a uint64) {
	cs.mux.Lock()
	defer cs.mux.Unlock()

	cs.redSum += r
	cs.greenSum += g
	cs.blueSum += b
	cs.alphaSum += a
}

func AverageColor(img image.Image) color.NRGBA {
	bounds := img.Bounds()
	cs := &colorSums{}

	var wg sync.WaitGroup
	wg.Add(bounds.Max.Y)
	for y := 0; y < bounds.Max.Y; y++ {
		go func(y int) {
			defer wg.Done()
			var reds, greens, blues, alphas uint64
			for x := 0; x < bounds.Max.X; x++ {
				redAP, greenAP, blueAP, alphaAP := img.At(x, y).RGBA() // alpha-premultiplied values
				reds += uint64(redAP)
				greens += uint64(greenAP)
				blues += uint64(blueAP)
				alphas += uint64(alphaAP)
			}
			cs.inc(reds, greens, blues, alphas)
		}(y)
	}
	wg.Wait()

	if cs.alphaSum == 0 {
		return color.NRGBA{0, 0, 0, 0}
	}

	red := uint8(math.RoundToEven(float64(cs.redSum*0xff) / float64(cs.alphaSum)))
	green := uint8(math.RoundToEven(float64(cs.greenSum*0xff) / float64(cs.alphaSum)))
	blue := uint8(math.RoundToEven(float64(cs.blueSum*0xff) / float64(cs.alphaSum)))
	maxPossibleSum := bounds.Max.X * bounds.Max.Y * 0xffff
	alpha := uint8(math.RoundToEven(float64(cs.alphaSum*0xff) / float64(maxPossibleSum)))

	return color.NRGBA{red, green, blue, alpha}
}
