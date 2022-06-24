package container

import (
	"fmt"
)

//

type KeyError[K comparable] struct {
	Key K
}

func (e KeyError[K]) String() string {
	return fmt.Sprintf("keyError{%v}", e.Key)
}

func (e KeyError[K]) Error() string {
	return e.String()
}
