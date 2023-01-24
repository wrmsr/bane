package lifecycles

type StateError struct {
	Current  State
	Expected StateMask
}
