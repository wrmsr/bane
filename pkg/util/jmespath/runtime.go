package jmespath

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/check"
	ju "github.com/wrmsr/bane/pkg/util/json"
	"github.com/wrmsr/bane/pkg/util/maps"
	bt "github.com/wrmsr/bane/pkg/util/types"
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

//

type SimpleRuntime struct{}

var _ Runtime[any] = SimpleRuntime{}

func (r SimpleRuntime) IsTruthy(obj any) bool {
	switch r.GetType(obj) {
	case NullType:
		return false
	case NumberType:
		return true
	case StringType:
		return obj.(string) != ""
	case BooleanType:
		return obj.(bool)
	case ArrayType:
		return len(obj.([]any)) > 0
	case ObjectType:
		return len(obj.(map[string]any)) > 0
	}
	panic(fmt.Errorf("type error: %v", obj))
}

func (r SimpleRuntime) GetType(obj any) ValueType {
	if obj == nil {
		return NullType
	}
	switch obj.(type) {
	case int64:
		return NumberType
	case float64:
		return NumberType
	case string:
		return StringType
	case bool:
		return BooleanType
	case []any:
		return ArrayType
	case map[string]any:
		return ObjectType
	}
	panic(fmt.Errorf("type error: %v", obj))
}

func (r SimpleRuntime) IsNull(obj any) bool {
	return r.GetType(obj) == NullType
}

func (r SimpleRuntime) CreateNull() any {
	return nil
}

func (r SimpleRuntime) asFloat(o any) float64 {
	switch o := o.(type) {
	case bool:
		return 1
	case int64:
		return float64(o)
	case float64:
		return o
	}
	panic(fmt.Errorf("type error: %v", o))
}

func (r SimpleRuntime) Compare(o CmpOp, left, right any) any {
	lty := r.GetType(left)
	rty := r.GetType(right)
	if lty != rty {
		return -1
	}
	switch lty {
	case NullType:
		return 0
	case BooleanType:
		return bt.BoolCmp(left.(bool), right.(bool))
	case NumberType:
		if bt.Is[int64](left) && bt.Is[int64](right) {
			return bt.OrderedCmp(left.(int64), right.(int64))
		}
		return bt.OrderedCmp(r.asFloat(left), r.asFloat(right))
	}
	panic(fmt.Errorf("type error: %v, %v", left, right))
}

func (r SimpleRuntime) CreateArray(items []any) any {
	return items
}

func (r SimpleRuntime) CreateObject(fields map[string]any) any {
	return fields
}

func (r SimpleRuntime) ToIterable(obj any) []any {
	switch obj := obj.(type) {
	case []any:
		return obj
	case map[string]any:
		return maps.Values(obj)
	}
	return nil
}

func (r SimpleRuntime) InvokeFunction(name string, args []any) any {
	panic("implement me")
}

func (r SimpleRuntime) CreateBool(value bool) any {
	return value
}

func (r SimpleRuntime) GetProperty(obj any, field string) any {
	switch obj := obj.(type) {
	case map[string]any:
		return obj
	}
	return nil
}

func (r SimpleRuntime) ParseStr(s string) any {
	return check.Must1(ju.UnmarshalAs[any]([]byte(s)))
}

func (r SimpleRuntime) CreateStr(val string) any {
	return val
}

func (r SimpleRuntime) GetNumVar(num int) any {
	return num
}

func (r SimpleRuntime) GetNameVar(name string) any {
	panic("implement me")
}
