package jmespath

type ValueType int8

const (
	InvalidType ValueType = iota
	NumberType
	StringType
	BooleanType
	ArrayType
	ObjectType
	NullType
)
