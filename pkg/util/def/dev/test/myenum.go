//go:build !nodev

package test

type MyEnum int8

const (
	InvalidMyEnum MyEnum = iota
	FooMyEnum
	BarMyEnum
)
