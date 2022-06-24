package tpch

import (
	"fmt"
	"testing"
)

func TestTpch(t *testing.T) {
	fmt.Println(getTextDists())
	fmt.Printf("%+v\n", Entities)
}
