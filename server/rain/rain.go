package rain

import (
	"math/rand"
	"time"
)

type Cloud struct {
	buf            []byte
	frameCount     float64
	Speed          float64
	RainfallAmount float64
}

func NewCloud() *Cloud {
	rand.Seed(time.Now().UnixNano())
	c := &Cloud{Speed: 0.3, frameCount: 0, RainfallAmount: 0.05}
	c.buf = make([]byte, 64)
	return c
}

func (c *Cloud) rain() []byte {
	for i := 0; i < 8; i++ {
		tmp := byte(0x00)
		for j := 0; j < 8; j++ {
			r := rand.Float64()
			tmp = tmp << 1
			if r < c.RainfallAmount {
				tmp = tmp | 0x01
			}
		}
		c.buf = c.buf[:len(c.buf)-1]
		c.buf, c.buf[0] = append(c.buf[:1], c.buf[0:]...), tmp
	}
	return c.buf[:]
}

func (c *Cloud) Generate(frame int) []byte {
	c.frameCount = c.frameCount + c.Speed
	if c.frameCount > 1 {
		c.frameCount = 0
		return c.rain()
	}
	return c.buf
}
