package containers

import (
	"fmt"
	"math/rand"
	"testing"
)

type IntComparer struct{}

func (i IntComparer) Compare(left, right int) int {
	return left - right
}

func TestTreapMap(t *testing.T) {
	rand.Seed(42)

	m := NewTreapMap[int, string](IntComparer{})
	fmt.Println("Count:", m.Count())

	m = m.Set(52, "hello")
	m = m.Set(53, "world")
	m = m.Set(52, "Hello")

	m.ForEach(func(k int, v string) {
		fmt.Println("[", k, "] =", v)
	})
	fmt.Println("Count:", m.Count())

	old := m.Set(500, "five hundred")
	m = m.Delete(53)

	fmt.Println(m.Get(53))
	fmt.Println(old.Get(53))
	fmt.Println(old.Get(52))
	fmt.Println(old.Get(500))

}
