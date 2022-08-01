package inject

import (
	"fmt"
	"strings"
)

//

type Error interface {
	isError()
}

//

type GenericError struct {
	Err error
}

func genericError(err error) GenericError {
	return GenericError{Err: err}
}

func genericErrorf(format string, a ...any) GenericError {
	return GenericError{Err: fmt.Errorf(format, a...)}
}

func (e GenericError) isError() {}

func (e GenericError) Error() string { return e.Err.Error() }
func (e GenericError) Unwrap() error { return e.Err }

//

type UnboundKeyError struct {
	Key Key
	Src any
}

func (e UnboundKeyError) isError() {}

func (e UnboundKeyError) Error() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("unbound key: %+v", e.Key))
	if e.Src != nil {
		sb.WriteString(fmt.Sprintf(" from src: %s", e.Src))
	}
	return sb.String()
}

//

type DuplicateBindingError struct {
	Key Key
}

func (e DuplicateBindingError) isError() {}

func (e DuplicateBindingError) Error() string {
	return fmt.Sprintf("duplicate binding: %+v", e.Key)
}
