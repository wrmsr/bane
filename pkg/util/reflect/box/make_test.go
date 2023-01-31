package box

import (
	"fmt"
	"testing"
)

func TestBox(t *testing.T) {
	b := BoxOf(420)
	fmt.Println(b)
}
