package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//var palette = []color.Color{color.White, color.Black}
var palette = []color.Color{color.RGBA{0x00, 0x00, 0x00, 0xff},

	color.RGBA{0xFF, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xFF, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xFF, 0xff},

	color.RGBA{0xFF, 0xFF, 0x00, 0xff},
	color.RGBA{0xFF, 0x00, 0xFF, 0xff},
	color.RGBA{0x00, 0xFF, 0xFF, 0xff}}

//var palette = []color.RGBA{0xRR, 0xGG, 0xBB, 0xff}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	lissajous(os.Stdout)
}

//var (
//	i uint = 0
//)

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100
		nframes = 64
		delay   = 8

		// image canvas covers [-size..+size]
		// number of animation frames
		// delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	// i = 0
	for i := 0; i < nframes; i++ {
		var i int = 0
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			rem := i / 50 % len(palette)
			i++
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rem))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
