package tower

//

type ContainerType interface {
	isContainerType()
}

type BaseContainerType struct {
	BaseType
}

func (t BaseContainerType) isContainerType() {}

//

type Container interface {
	isContainer()
}

type BaseContainer struct {
	BaseValue
}

func (v BaseContainer) isContainer() {}
