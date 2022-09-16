package arm64

import (
	"fmt"
	"regexp"
	"testing"
)

func TestDa(t *testing.T) {
	otp := regexp.MustCompile("^(.+)_([0-9%*])$")

	for k, _ := range map_op {
		m := otp.FindStringSubmatch(k)
		fmt.Printf("%s: %v\n", k, m[1])
	}
}
