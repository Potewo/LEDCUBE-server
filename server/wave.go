package main

type Waver struct {
	buf [8][8]byte
}

func (w *Waver) Generate(frame int) [8][8]byte {
	for i, bs := range w.buf {
		for j := range bs {
			w.buf[i][j] = 0x00
		}
	}
	return w.buf
}
