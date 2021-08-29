package wave

import (
	"math"
)

type Waver struct {
	buf [64]byte
	freq float64
}

func NewWaver() *Waver{
	return &Waver{freq:10}
}

func (w Waver) wave(seed int) int {
	return int(math.Round(4 * math.Sin(float64(seed) * 2 * math.Pi / w.freq))) + 3
}

func (w *Waver) Generate(frame int) []byte {
	for i := 0; i < 8; i++ {
		h := w.wave(i + frame)
		// fmt.Println(h)
		for j := 0; j < 8; j++ {
			if j == h {
				w.buf[i*8+j] = 0xff
			} else {
				w.buf[i*8+j] = 0x00
			}
		}
	}
	return w.buf[:]
	// 配列からスライスへの変換
}
