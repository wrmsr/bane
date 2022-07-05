package funcs

import "fmt"

//

type PanicErr struct {
	Value any
}

func (e PanicErr) Error() string {
	return fmt.Sprintf("%v", e.Value)
}

func (e PanicErr) String() string {
	return fmt.Sprintf("PanicErr{%v}", e.Value)
}

//

func Recover(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicErr{Value: r}
		}
	}()
	fn()
	return
}
