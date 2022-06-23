package funcs

import (
	"fmt"
	"reflect"
	"text/template"
)

var _ = Std.AddAll(template.FuncMap{
	"typeIs":     TypeIs,
	"typeIsLike": TypeIsLike,
	"typeOf":     TypeOf,
	"kindIs":     KindIs,
	"kindOf":     KindOf,
})

func TypeIs(target string, src interface{}) bool {
	return target == TypeOf(src)
}

func TypeIsLike(target string, src interface{}) bool {
	t := TypeOf(src)
	return target == t || "*"+target == t
}

func TypeOf(src interface{}) string {
	return fmt.Sprintf("%T", src)
}

func KindIs(target string, src interface{}) bool {
	return target == KindOf(src)
}

func KindOf(src interface{}) string {
	return reflect.ValueOf(src).Kind().String()
}
