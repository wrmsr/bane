package box

//

type InterfaceType struct {
	type_
}

//

type Interface struct {
	value
}

func (v Interface) Type() Type { return BoxType(v.v.Type()) }
