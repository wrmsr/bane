package lifecycles

type Handler = func(st State) error
type Callback = func(obj any, st State)

type Lifecycle interface {
	State() State
	AddCallback(cb Callback)

	Construct() error
	Start() error
	Stop() error
	Destroy() error
}
