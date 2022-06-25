package unsafe

import (
	"fmt"
	"testing"
)

func TestAsm(t *testing.T) {
	fmt.Println(__addints(3, 4))
}
