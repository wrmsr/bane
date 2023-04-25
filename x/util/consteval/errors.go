package consteval

import (
	"fmt"
	"strings"
)

//

type NameError struct {
	N string
}

func (e NameError) Error() string {
	return fmt.Sprintf("unresolved name: %s", e.N)
}

//

type TypeError struct {
	S []string
}

func (e TypeError) Error() string {
	return fmt.Sprintf("type error: %s", strings.Join(e.S, ", "))
}
