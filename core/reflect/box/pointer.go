package box

//

type PointerType struct {
	type_
}

//

type Pointer struct {
	value
}

func (v Pointer) Type() Type { return BoxType(v.v.Type()) }
