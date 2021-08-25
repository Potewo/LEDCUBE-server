package wave

import (
	"fmt"
	"testing"
)

func TestGenerateSuccess(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", Generate(i))
	}
}
