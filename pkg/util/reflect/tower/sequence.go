package tower

//

type SequenceType interface {
	isSequenceType()
}

type sequenceType struct {
	containerType
}

func (t sequenceType) isSequenceType() {}

//

type Sequence interface {
	isSequence()
}

type sequence struct {
	container
}

func (v sequence) isSequence() {}
