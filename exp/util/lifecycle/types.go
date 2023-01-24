package lifecycles

//

type Lifecycle struct {
	Fn func(State) error
}

//

type Callback func(*Lifecycle, State)

type Controller interface {
	State() State
	AddCallback(Callback)

	Construct() error
	Start() error
	Stop() error
	Destroy() error
}
