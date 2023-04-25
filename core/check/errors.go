package check

import "fmt"

type CheckError struct {
	Message string
	Args    []any
	Details []any
}

func checkError(msg string, args []any, details ...any) CheckError {
	return CheckError{Message: msg, Args: args, Details: details}
}

func (e CheckError) String() string {
	s := e.Message
	if len(e.Args) > 0 {
		s = fmt.Sprintf("%s: %v", s, e.Args)
	}
	return s
}

func (e CheckError) Error() string {
	return e.String()
}
