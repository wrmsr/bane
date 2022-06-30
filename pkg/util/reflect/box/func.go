package box

//

type FuncType struct {
	type_
}

//

type Func struct {
	value
}

func (v Func) Type() Type { return BoxType(v.v.Type()) }
