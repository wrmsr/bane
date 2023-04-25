package sql

//

type NotFoundError struct{}

func (e NotFoundError) Error() string { return "not found" }

//

type MultipleFoundError struct{}

func (e MultipleFoundError) Error() string { return "multiple found" }

//

type NotEnoughValuesError struct{}

func (e NotEnoughValuesError) Error() string { return "not enough values" }

//

type TooManyValuesError struct{}

func (e TooManyValuesError) Error() string { return "too many values" }
