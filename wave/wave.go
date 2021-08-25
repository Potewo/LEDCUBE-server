package wave

import (
	"math"
)

var i = 0
var freq = 10.0
func Generate(i int) int {
	return int(math.Round(math.Sin(float64(i) * math.Pi * 2/ freq) * 4))
}
