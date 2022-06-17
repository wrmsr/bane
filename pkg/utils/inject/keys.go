package inject

import "reflect"

type Key struct {
	ty  reflect.Type
	arr bool
	tag any
}
