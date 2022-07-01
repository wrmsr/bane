package reflect

import "reflect"

var (
	_invalidValue     reflect.Value
	_trueValue        = reflect.ValueOf(true)
	_falseValue       = reflect.ValueOf(false)
	_zeroValue        = reflect.ValueOf(0)
	_emptyStringValue = reflect.ValueOf("")
)

func Invalid() reflect.Value     { return _invalidValue }
func True() reflect.Value        { return _trueValue }
func False() reflect.Value       { return _falseValue }
func Zero() reflect.Value        { return _zeroValue }
func EmptyString() reflect.Value { return _emptyStringValue }
