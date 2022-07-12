package iterators

import (
	"reflect"

	"github.com/wrmsr/bane/pkg/util/check"
)

type reflectHelperIterableMethods struct {
	iterateMth reflect.Value
}

type reflectHelperIteratorMethods struct {
	hasNextMeth, nextMeth reflect.Value
}

type ReflectHelper struct {
	iterableMeths map[reflect.Type]reflectHelperIterableMethods
	iteratorMeths map[reflect.Type]reflectHelperIteratorMethods
}

func NewReflectHelper() ReflectHelper {
	return ReflectHelper{
		iterableMeths: make(map[reflect.Type]reflectHelperIterableMethods),
		iteratorMeths: make(map[reflect.Type]reflectHelperIteratorMethods),
	}
}

func (h ReflectHelper) ForEach(iterableVal reflect.Value, fn func(v reflect.Value) bool) bool {
	iterableTy := iterableVal.Type()
	iterableMeths, ok := h.iterableMeths[iterableTy]
	if !ok {
		iterableMeths = reflectHelperIterableMethods{
			iterateMth: check.Ok1(iterableTy.MethodByName("Iterate")).Func,
		}
		h.iterableMeths[iterableTy] = iterableMeths
	}

	iteratorVal := check.Single(iterableMeths.iterateMth.Call([]reflect.Value{iterableVal}))
	if iteratorVal.Kind() == reflect.Interface {
		iteratorVal = iteratorVal.Elem()
	}
	iteratorTy := iteratorVal.Type()

	iteratorMeths, ok := h.iteratorMeths[iteratorTy]
	if !ok {
		iteratorMeths = reflectHelperIteratorMethods{
			nextMeth:    check.Ok1(iteratorTy.MethodByName("Next")).Func,
			hasNextMeth: check.Ok1(iteratorTy.MethodByName("HasNext")).Func,
		}
		h.iteratorMeths[iteratorTy] = iteratorMeths
	}

	for check.Single(iteratorMeths.hasNextMeth.Call([]reflect.Value{iteratorVal})).Bool() {
		elemVal := check.Single(iteratorMeths.nextMeth.Call([]reflect.Value{iteratorVal}))
		if !fn(elemVal) {
			return false
		}
	}
	return true
}
