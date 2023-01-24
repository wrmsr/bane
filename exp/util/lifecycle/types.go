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

type Controller interface {
	Construct() error
	Start() error
	Stop() error
	Destroy() error
}

//

type Whence int8

const (
	Before Whence = iota
	After
)

type Callback func(*Lifecycle, Whence, State)
