//go:build !nodev

package test

import "github.com/wrmsr/bane/core/def"

var _ = def.Enum[MyEnum]()

type MyEnum int8

const (
	InvalidMyEnum MyEnum = iota
	FooMyEnum
	BarMyEnum
)
