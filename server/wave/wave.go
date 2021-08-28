package wave

import (
	"fmt"
	"math"
)

type Waver struct {
	buf [8][8]byte
	freq float64
}

func NewWaver() *Waver{
	return &Waver{freq:10}
}

func (w Waver) wave(seed int) int {
	s := float64(seed) * 2 * math.Pi / w.freq
	fmt.Println(s)
	return int(math.Round(4 * math.Sin(float64(seed) * 2 * math.Pi / w.freq))) + 3
}

func (w *Waver) Generate(frame int) [8][8]byte {
	for i, bs := range w.buf {
		h := w.wave(i + frame)
		// fmt.Println(h)
		for j := range bs {
			if j == h {
				w.buf[i][j] = 0xff
			} else {
				w.buf[i][j] = 0x00
			}
		}
	}
	return w.buf
}
