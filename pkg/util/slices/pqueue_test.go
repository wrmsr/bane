package slices

import (
	"fmt"
	"testing"
)

func TestPq(t *testing.T) {
	var pq PriorityQueue[string]
	pq.Init([]PqInitItem[string]{
		{Value: "ten", Priority: 10},
		{Value: "one", Priority: 1},
		{Value: "two", Priority: 2},
		{Value: "zero", Priority: 0},
		{Value: "twenty", Priority: 20},
	})
	for _, item := range pq.s {
		fmt.Println(item)
	}
	fmt.Println()
	for pq.Len() > 0 {
		fmt.Println(pq.Pop())
	}
}
