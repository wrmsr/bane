package box

//

type UnsafePointerType struct {
	type_
}

//

type UnsafePointer struct {
	value
}

func (v UnsafePointer) Type() Type { return BoxType(v.v.Type()) }
