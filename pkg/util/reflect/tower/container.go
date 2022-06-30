package tower

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
