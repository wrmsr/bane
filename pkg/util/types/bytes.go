package types

//

type Bytes struct {
	s []byte
}

func BytesOf(s []byte) Bytes {
	return Bytes{s: CloneSlice(s)}
}

func UnsafeBytesOf(s []byte) Bytes {
	return Bytes{s: s}
}

func (b Bytes) Len() int       { return len(b.s) }
func (b Bytes) Get(i int) byte { return b.s[i] }

func (b Bytes) Slice() []byte { return CloneSlice(b.s) }

//

var _ HashEq[Bytes] = Bytes{}

func (b Bytes) Hash() uintptr {
	return HashBytes(b.s)
}

func (b Bytes) Equals(o Bytes) bool {
	if len(b.s) != len(o.s) {
		return false
	}
	for i, l := range o.s {
		if l != o.s[i] {
			return false
		}
	}
	return true
}

//

var _ Comparer[Bytes] = Bytes{}

func (b Bytes) Compare(o Bytes) CmpResult {
	bl := len(b.s)
	ol := len(o.s)
	mn := Max(bl, ol)
	for i := 0; i < mn; i++ {
		if i >= bl {
			return CmpLesser
		}
		if i >= ol {
			return CmpGreater
		}
		bb := b.s[i]
		ob := o.s[i]
		if bb < ob {
			return CmpLesser
		}
		if ob < bb {
			return CmpGreater
		}
	}
	return CmpEqual
}

//

var _ Iterable[byte] = Bytes{}

func (b Bytes) Iterate() Iterator[byte] {
	panic("implement me")
}

type bytesIterator struct {
	s []byte
	i int
}

var _ Iterator[byte] = &bytesIterator{}

func (i *bytesIterator) Iterate() Iterator[byte] {
	return i
}

func (i *bytesIterator) HasNext() bool {
	return i.i < len(i.s)
}

func (i *bytesIterator) Next() byte {
	if i.i >= len(i.s) {
		panic(IteratorExhaustedError{})
	}
	b := i.s[i.i]
	i.i++
	return b
}

//

var _ Traversable[byte] = Bytes{}

func (b Bytes) ForEach(fn func(v byte) bool) bool {
	for i := 0; i < len(b.s); i++ {
		if !fn(b.s[i]) {
			return false
		}
	}
	return true
}
