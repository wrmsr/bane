package inject

import "fmt"

//

type Error interface {
	isError()
}

//

type UnboundKeyError struct {
	Key Key
}

func (e UnboundKeyError) isError() {}

func (e UnboundKeyError) Error() string {
	return fmt.Sprintf("unbound key: %+v", e.Key)
}

//

type DuplicateBindingError struct {
	Key Key
}

func (e DuplicateBindingError) isError() {}

func (e DuplicateBindingError) Error() string {
	return fmt.Sprintf("duplicate binding: %+v", e.Key)
}
