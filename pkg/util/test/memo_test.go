package test

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

//

type MemoStruct struct {
	_MemoStruct

	x int

	numFoo int
}

func (m *MemoStruct) _Foo() int {
	m.numFoo++
	return 420 + m.x
}

func (m *MemoStruct) _Bar() int {
	return m.Foo() + 1
}

//

type _MemoStruct struct {
	__MemoStruct *MemoStruct

	__Foo opt.Optional[int]
	__Bar opt.Optional[int]
}

func (m *_MemoStruct) Foo() int {
	if !m.__Foo.Present() {
		m.__Foo = opt.Just(m.__MemoStruct._Foo())
	}
	return m.__Foo.Value()
}

func (m *_MemoStruct) Bar() int {
	if !m.__Bar.Present() {
		m.__Bar = opt.Just(m.__MemoStruct._Bar())
	}
	return m.__Bar.Value()
}

//

func TestMemoMockup(t *testing.T) {
	m := MemoStruct{
		x: 10,
	}
	m._MemoStruct.__MemoStruct = &m

	tu.AssertEqual(t, m.numFoo, 0)
	tu.AssertEqual(t, m.Foo(), 430)
	tu.AssertEqual(t, m.numFoo, 1)

	tu.AssertEqual(t, m.Bar(), 431)
	tu.AssertEqual(t, m.numFoo, 1)
}
