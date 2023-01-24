package lifecycles

//

type LifecycleFn = func() error

type Lifecycle struct {
	Construct LifecycleFn
	Start     LifecycleFn
	Stop      LifecycleFn
	Destroy   LifecycleFn
}

//

type Whence int8

const (
	Before Whence = iota
	After
)

type Callback func(*Lifecycle, Whence, State)
