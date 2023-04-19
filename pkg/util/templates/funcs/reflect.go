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

func TypeIs(target string, src any) bool {
	return target == TypeOf(src)
}

func TypeIsLike(target string, src any) bool {
	t := TypeOf(src)
	return target == t || "*"+target == t
}

func TypeOf(src any) string {
	return fmt.Sprintf("%T", src)
}

func KindIs(target string, src any) bool {
	return target == KindOf(src)
}

func KindOf(src any) string {
	return reflect.ValueOf(src).Kind().String()
}
