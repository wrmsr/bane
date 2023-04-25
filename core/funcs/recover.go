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

func Recover1[T any](fn func() T) (r T, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicErr{Value: r}
		}
	}()
	r = fn()
	return
}

//

func Recovering(fn func()) func() error {
	return func() error {
		return Recover(fn)
	}
}

func Recovering1[T any](fn func() T) func() (T, error) {
	return func() (T, error) {
		return Recover1[T](fn)
	}
}
