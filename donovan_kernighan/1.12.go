package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x42, 0x8f, 0x28, 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	var cyclesParam int
	http.HandleFunc("/", func(resp http.ResponseWriter, r *http.Request) {
		val, err := url.ParseQuery(r.URL.RawQuery)
		for k, v := range val {
			if k == "cycles" {
				cyclesParam, err = strconv.Atoi(v[0])
				must(err)
			}
		}

		must(err)
		lissajous(resp, cyclesParam)
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func lissajous(out io.Writer, param int) {
	var cyc float64
	const (
		cycles  = 5     // кол-во полных колебаний
		res     = 0.001 // угловое разрешение
		size    = 100   // канва изображения охватывает [size..+size]
		nframes = 64    // Кол-во кадров анимации
		delay   = 8     // задержка между кадрами (единица - 10мс)
	)
	if param > 0 {
		cyc = float64(param)
	} else {
		cyc = cycles
	}
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cyc*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
