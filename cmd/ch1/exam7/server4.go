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
)

var palette = []color.Color{
	color.White,
	color.Black,
}

type PaletteColor int

const (
	PaletteColorWhite PaletteColor = 0 + iota // iota는 enumerator (행이 바뀔 때마다 값이 차례로 증가)
	PaletteColorBlack
)

// go run server4.go
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}) // 각 요청은 핸들러를 호출
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // x 진동자의 회전수
		res     = 0.001 // 회전각
		size    = 100   // 이미지 캔버스 크기
		nframes = 64    // 애니메이션 프레임 수
		delay   = 8     // 10ms 단위의 프레임 간 지연
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes} // gif.GIF는 구조체 (서로 다른 타입들을 하나의 객체로 묶어 단일 객체로 취급
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(PaletteColorBlack))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
