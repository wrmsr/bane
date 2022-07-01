package reflect

import (
	"reflect"
	"time"
)

var (
	_boolType       = TypeOf[bool]()
	_intType        = TypeOf[int]()
	_int8Type       = TypeOf[int8]()
	_int16Type      = TypeOf[int16]()
	_int32Type      = TypeOf[int32]()
	_int64Type      = TypeOf[int64]()
	_uintType       = TypeOf[uint]()
	_uint8Type      = TypeOf[uint8]()
	_uint16Type     = TypeOf[uint16]()
	_uint32Type     = TypeOf[uint32]()
	_uint64Type     = TypeOf[uint64]()
	_uintptrType    = TypeOf[uintptr]()
	_float32Type    = TypeOf[float32]()
	_float64Type    = TypeOf[float64]()
	_complex64Type  = TypeOf[complex64]()
	_complex128Type = TypeOf[complex128]()
	_stringType     = TypeOf[string]()
	_anyType        = TypeOf[any]()
	_bytesType      = TypeOf[[]byte]()
	_timeType       = TypeOf[time.Time]()
	_durationType   = TypeOf[time.Duration]()
)

func Bool() reflect.Type       { return _boolType }
func Int() reflect.Type        { return _intType }
func Int8() reflect.Type       { return _int8Type }
func Int16() reflect.Type      { return _int16Type }
func Int32() reflect.Type      { return _int32Type }
func Int64() reflect.Type      { return _int64Type }
func Uint() reflect.Type       { return _uintType }
func Uint8() reflect.Type      { return _uint8Type }
func Uint16() reflect.Type     { return _uint16Type }
func Uint32() reflect.Type     { return _uint32Type }
func Uint64() reflect.Type     { return _uint64Type }
func Uintptr() reflect.Type    { return _uintptrType }
func Float32() reflect.Type    { return _float32Type }
func Float64() reflect.Type    { return _float64Type }
func Complex64() reflect.Type  { return _complex64Type }
func Complex128() reflect.Type { return _complex128Type }
func String() reflect.Type     { return _stringType }
func Any() reflect.Type        { return _anyType }
func Bytes() reflect.Type      { return _bytesType }
func Time() reflect.Type       { return _timeType }
func Duration() reflect.Type   { return _durationType }
