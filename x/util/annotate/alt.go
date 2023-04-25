package annotate

/*
//

func Type[T any]() any {
	return nil
}

type Annotator interface {
	Field(n string, anns ...any) any
}

func Annotate(obj any, anns ...any) Annotator {
	return nil
}

//

type someStruct struct {
	I int
}

func someFunc() {}

type someAnn struct{}

var _ = Annotate(Type[someStruct](),
	someAnn{},
).Field("I",
	nil,
)

var _ = Annotate(someFunc, someAnn{})
*/
