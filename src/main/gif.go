package main

import (
	"bufio"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.RGBA{0x00, 0x80, 0x00, 0xff}, color.Black}

const (
	whiteIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 150   // image canvas covers [-size..+size]
		nframes = 8     // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			ii := i % 2
			if ii == 0 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					blackIndex)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					greenIndex)
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	fo, _ := os.Create("output.gif")
	fw := bufio.NewWriter(fo)
	gif.EncodeAll(fw, &anim) // NOTE: ignoring encoding errors
	fw.Flush()
	fo.Close()

}
