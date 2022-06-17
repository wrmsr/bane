package inject

import "fmt"

//

type Error interface {
	isError()
}

//

type DuplicateBindingError struct {
	Key Key
}

func (e DuplicateBindingError) isError() {}

func (e DuplicateBindingError) Error() string {
	return fmt.Sprintf("duplicate binding: %+v", e.Key)
}
