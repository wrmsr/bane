package lifecycles

//

type LifecycleFn = func() error

type Lifecycle struct {
	Construct,
	Start,
	Stop,
	Destroy LifecycleFn
}

//

type Whence int8

const (
	Before Whence = iota
	After
)

type Callback func(*Lifecycle, Whence, State)
