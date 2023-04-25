package reflect

import (
	"fmt"
	"testing"

	tu "github.com/wrmsr/bane/core/dev/testing"
)

type IA interface {
	A()
}

type IB interface {
	B()
}

type SAB struct {
	i int
	s string
}

func (sab SAB) A() {
	fmt.Println(sab.i)
}

func (sab SAB) B() {
	fmt.Println(sab.s)
}

func TestAddressOf(t *testing.T) {
	sab := SAB{i: 420, s: "four twenty"}
	var a0 IA = &sab
	var a1 IA = &sab
	var b0 IB = &sab
	var b1 IB = &sab
	addr := AddressOf(&sab)
	tu.AssertEqual(t, addr, AddressOf(a0))
	tu.AssertEqual(t, addr, AddressOf(a1))
	tu.AssertEqual(t, addr, AddressOf(b0))
	tu.AssertEqual(t, addr, AddressOf(b1))
}
