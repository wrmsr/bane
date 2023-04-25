package box

//

type ChanType struct {
	type_
}

//

type Chan struct {
	value
}

func (v Chan) Type() Type { return BoxType(v.v.Type()) }
