package wave

import (
	"math"
)

type Waver struct {
	buf [64]byte
	Freq float64
	Speed float64
}

func NewWaver() *Waver{
	return &Waver{Freq:15, Speed: 0.3}
}

func (w Waver) wave(seed float64) int {
	return int(math.Round(4 * math.Sin(seed * 2 * math.Pi / w.Freq))) + 3
}

func (w *Waver) Generate(frame int) []byte {
	for i := 0; i < 8; i++ {
		h := w.wave(float64(i) + (float64(frame) * w.Speed))
		// fmt.Println(h)
		for j := 0; j < 8; j++ {
			if j == h {
				w.buf[j*8+i] = 0xff
			} else {
				w.buf[j*8+i] = 0x00
			}
		}
	}
	return w.buf[:]
	// 配列からスライスへの変換
}
