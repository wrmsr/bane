package tower

//

type ContainerType interface {
	isContainerType()
}

type containerType struct {
	type_
}

func (t containerType) isContainerType() {}

//

type Container interface {
	isContainer()
}

type container struct {
	value
}

func (v container) isContainer() {}
