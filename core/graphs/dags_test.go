package graphs

import (
	"fmt"
	"testing"

	ctr "github.com/wrmsr/bane/core/container"
)

func TestDags(t *testing.T) {
	m := map[int][]int{
		0: {1, 2},
		1: {2},
		2: nil,
		3: {4},
		4: nil,
	}

	d := NewDag(SliceEdges(m))
	fmt.Println(d.All())
	fmt.Println(d.OutputSetsByOutput())

	sd := d.Subdag(ctr.NewStdSetOf(1), nil)
	fmt.Println(sd.Inputs())
	fmt.Println(sd.Outputs())
	fmt.Println(sd.OutputInputs())
	fmt.Println(sd.All())
}
