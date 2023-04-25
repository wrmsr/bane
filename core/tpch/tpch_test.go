package tpch

import (
	"fmt"
	"testing"
)

func TestTpch(t *testing.T) {
	regionGen := generateRegions(nil, nil).Iterate()
	for i := 0; i < 3; i++ {
		fmt.Println(regionGen.Next())
	}
}
