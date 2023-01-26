package lifecycles

type Handler = func(State) error
type Callback = func(State)

type Lifecycle interface {
	State() State
	AddCallback(Callback)

	Construct() error
	Start() error
	Stop() error
	Destroy() error
}
