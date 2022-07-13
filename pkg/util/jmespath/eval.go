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

type Runtime[T any] interface {
	IsTruthy(obj T) bool
	GetType(obj T) ValueType
	IsNull(obj T) bool
	CreateNull() T
	Compare(o CmpOp, left, right T) T
	CreateArray(items []T) T
	CreateObject(fields map[string]T) T
	ToIterable(obj T) []T
	InvokeFunction(name string, args []any) T
	CreateBool(value bool) T
	GetProperty(obj T, field string) T
	ParseStr(s string) T
	CreateStr(val string) T
	GetNumVar(num int) T
	GetNameVar(name string) T
}
