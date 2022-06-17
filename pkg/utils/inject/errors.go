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

type UnboundKeyError struct {
	Key Key
	Src any
}

func (e UnboundKeyError) isError() {}

func (e UnboundKeyError) Error() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("unbound key: %+v", e.Key))
	if e.Src != nil {
		sb.WriteString(fmt.Sprintf(" from src: %v", e.Src))
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
