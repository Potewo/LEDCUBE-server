package plane

import "math"

type Plane struct {
	buf [64]byte
	Freq float64
}

func NewPlane() *Plane{
	return &Plane{Freq: 3}
}

func (p Plane) planeZ(frame int) []byte {
	for i := 0; i < 8; i++ {
		if int(math.Round(float64(frame) / p.Freq)) % 8 == i {
			for j := 0; j < 8; j++ {
				p.buf[8*i+j] = 0xff
			}
		} else {
			for j := 0; j < 8; j++ {
				p.buf[8*i+j] = 0x00
			}
		}
	}
	return p.buf[:]
}

func (p Plane) planeX(frame int) []byte {
	for i := 0; i < 8; i++ {
		if int (math.Round(float64(frame) / p.Freq)) % 8 == i {
			for j := 0; j < 8; j++ {
				p.buf[8*j+i] = 0xff
			}
		} else {
			for j := 0; j < 8; j++ {
				p.buf[8*j+i] = 0x00
			}
		}
	}
	return p.buf[:]
}

func (p Plane) planeY(frame int) []byte {
	index := int(math.Round(float64(frame) / p.Freq)) % 8
	b := byte(0x01 << index)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			p.buf[8*i+j] = b
		}
	}
	return p.buf[:]
}

func (p Plane) Generate(frame int) []byte {
	return p.planeY(frame)
}
