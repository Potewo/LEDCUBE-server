package wave

import (
	"fmt"
	"testing"
)

func TestGenerateSuccess(t *testing.T) {
	waver := NewWaver()
	for f := 0; f < 30; f++ {
		result := waver.Generate(1)
		fmt.Println(result)
	}
}
