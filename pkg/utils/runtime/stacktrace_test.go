package runtime

import (
	"fmt"
	"testing"
)

func TestGetStackTrace(t *testing.T) {
	sfs := GetStackTrace(2, 0)
	fmt.Println(sfs)
	fmt.Println(ParseName(sfs[0].Name))
}
