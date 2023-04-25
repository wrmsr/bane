package types

import "testing"

type anyEqer struct{}

func (e anyEqer) Equals(o any) bool { return false }
func (e anyEqer) Hash() uintptr     { return 0 }

var _ HashEq[any] = anyEqer{}

func TestHashEq(t *testing.T) {

}
