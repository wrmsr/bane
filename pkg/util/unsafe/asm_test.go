package unsafe

import (
	"fmt"
	"testing"
)

func TestAsm(t *testing.T) {
	fmt.Printf("%x\n", __addints(3, 4))
}
