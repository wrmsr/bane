package inject

type Binding struct {
	key      Key
	provider Provider
}

type Bindings struct {
	bs []Binding
}
