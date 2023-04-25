package box

//

type ContainerType interface {
	Type
	isContainerType()
}

type containerType struct {
	type_
}

func (t containerType) isContainerType() {}

//

type Container interface {
	Value
	isContainer()
}

type container struct {
	value
}

func (v container) isContainer() {}

func (v container) Type() Type { return BoxType(v.v.Type()) }
