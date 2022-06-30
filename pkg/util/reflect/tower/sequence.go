package tower

//

type SequenceType interface {
	isSequenceType()
}

type BaseSequenceType struct {
	BaseContainerType
}

func (t BaseSequenceType) isSequenceType() {}

//

type Sequence interface {
	isSequence()
}

type BaseSequence struct {
	BaseContainer
}

func (v BaseSequence) isSequence() {}
