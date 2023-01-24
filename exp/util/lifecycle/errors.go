package lifecycles

import "fmt"

type StateError struct {
	Current  State
	Expected StateMask
}

func (e StateError) Error() string {
	return fmt.Sprintf("state error: current %v, expected %v", e.Current, e.Expected)
}
