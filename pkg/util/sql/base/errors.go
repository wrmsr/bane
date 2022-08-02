package base

import "fmt"

//

type DuplicateColumnError struct {
	Name string
}

func (e DuplicateColumnError) Error() string { return fmt.Sprintf("duplicate column: %s", e.Name) }
