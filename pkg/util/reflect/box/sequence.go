package box

//

type SequenceType interface {
	ContainerType
	isSequenceType()
}

type sequenceType struct {
	containerType
}

func (t sequenceType) isSequenceType() {}

//

type Sequence interface {
	Container
	isSequence()
}

type sequence struct {
	container
}

func (v sequence) isSequence() {}
