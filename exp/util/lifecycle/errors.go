package lifecycles

import "fmt"

//

type StateError struct {
	Current  State
	Expected StateMask
}

func (e StateError) Error() string {
	return fmt.Sprintf("state error: current %v, expected %v", e.Current, e.Expected)
}

//

type DependencyNotFoundError struct {
	Dependent,
	Dependency any
}

func (e DependencyNotFoundError) Error() string {
	return fmt.Sprintf("dependency not found: %v -> %v", e.Dependent, e.Dependency)
}

//

type ObjectNotRegisteredError struct {
	Obj any
}

func (e ObjectNotRegisteredError) Error() string {
	return fmt.Sprintf("object not registered: %v", e.Obj)
}

//

type ObjectAlreadyRegisteredError struct {
	Obj any
}

func (e ObjectAlreadyRegisteredError) Error() string {
	return fmt.Sprintf("object already registered: %v", e.Obj)
}
