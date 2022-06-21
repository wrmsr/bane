package container

type Mutable interface {
	isMutable()
}

type Ordered interface {
	isOrdered()
}

type Persistent interface {
	isPersistent()
}

type Sorted interface {
	isSorted()
}
