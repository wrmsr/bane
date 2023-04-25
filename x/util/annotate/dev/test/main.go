//go:build !nodev

package main

import "github.com/wrmsr/bane/x/util/annotate"

type SomeAnn struct {
	S string
}

type SomeStruct struct {
	S string
	I int
}

func (s SomeStruct) SomeMethod() {}

var _ = annotate.On[SomeStruct](SomeAnn{}).
	Field("S", SomeAnn{}).
	Method("SomeMethod", SomeAnn{})

func main() {
}
