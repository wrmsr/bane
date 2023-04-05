// Copyright 2009 The Go Authors. All rights reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package binary

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
	"reflect"
)

//

type coder struct {
	order  binary.ByteOrder
	buf    []byte
	offset int
}

type Decoder coder
type Encoder coder

func (d *Decoder) Bool() bool {
	x := d.buf[d.offset]
	d.offset++
	return x != 0
}

func (e *Encoder) Bool(x bool) {
	if x {
		e.buf[e.offset] = 1
	} else {
		e.buf[e.offset] = 0
	}
	e.offset++
}

func (d *Decoder) Uint8() uint8 {
	x := d.buf[d.offset]
	d.offset++
	return x
}

func (e *Encoder) Uint8(x uint8) {
	e.buf[e.offset] = x
	e.offset++
}

func (d *Decoder) Uint16() uint16 {
	x := d.order.Uint16(d.buf[d.offset : d.offset+2])
	d.offset += 2
	return x
}

func (e *Encoder) Uint16(x uint16) {
	e.order.PutUint16(e.buf[e.offset:e.offset+2], x)
	e.offset += 2
}

func (d *Decoder) Uint32() uint32 {
	x := d.order.Uint32(d.buf[d.offset : d.offset+4])
	d.offset += 4
	return x
}

func (e *Encoder) Uint32(x uint32) {
	e.order.PutUint32(e.buf[e.offset:e.offset+4], x)
	e.offset += 4
}

func (d *Decoder) Uint64() uint64 {
	x := d.order.Uint64(d.buf[d.offset : d.offset+8])
	d.offset += 8
	return x
}

func (e *Encoder) Uint64(x uint64) {
	e.order.PutUint64(e.buf[e.offset:e.offset+8], x)
	e.offset += 8
}

func (d *Decoder) Int8() int8 { return int8(d.Uint8()) }

func (e *Encoder) Int8(x int8) { e.Uint8(uint8(x)) }

func (d *Decoder) Int16() int16 { return int16(d.Uint16()) }

func (e *Encoder) Int16(x int16) { e.Uint16(uint16(x)) }

func (d *Decoder) Int32() int32 { return int32(d.Uint32()) }

func (e *Encoder) Int32(x int32) { e.Uint32(uint32(x)) }

func (d *Decoder) Int64() int64 { return int64(d.Uint64()) }

func (e *Encoder) Int64(x int64) { e.Uint64(uint64(x)) }

func (d *Decoder) Value(v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		l := v.Len()
		for i := 0; i < l; i++ {
			d.Value(v.Index(i))
		}

	case reflect.Struct:
		t := v.Type()
		l := v.NumField()
		for i := 0; i < l; i++ {
			// Note: Calling v.CanSet() below is an optimization.
			// It would be sufficient to check the field name,
			// but creating the StructField info for each field is
			// costly (run "go test -bench=ReadStruct" and compare
			// results when making changes to this code).
			if v := v.Field(i); v.CanSet() || t.Field(i).Name != "_" {
				d.Value(v)
			} else {
				d.Skip(v)
			}
		}

	case reflect.Slice:
		l := v.Len()
		for i := 0; i < l; i++ {
			d.Value(v.Index(i))
		}

	case reflect.Bool:
		v.SetBool(d.Bool())

	case reflect.Int8:
		v.SetInt(int64(d.Int8()))
	case reflect.Int16:
		v.SetInt(int64(d.Int16()))
	case reflect.Int32:
		v.SetInt(int64(d.Int32()))
	case reflect.Int64:
		v.SetInt(d.Int64())

	case reflect.Uint8:
		v.SetUint(uint64(d.Uint8()))
	case reflect.Uint16:
		v.SetUint(uint64(d.Uint16()))
	case reflect.Uint32:
		v.SetUint(uint64(d.Uint32()))
	case reflect.Uint64:
		v.SetUint(d.Uint64())

	case reflect.Float32:
		v.SetFloat(float64(math.Float32frombits(d.Uint32())))
	case reflect.Float64:
		v.SetFloat(math.Float64frombits(d.Uint64()))

	case reflect.Complex64:
		v.SetComplex(complex(
			float64(math.Float32frombits(d.Uint32())),
			float64(math.Float32frombits(d.Uint32())),
		))
	case reflect.Complex128:
		v.SetComplex(complex(
			math.Float64frombits(d.Uint64()),
			math.Float64frombits(d.Uint64()),
		))
	}
}

func (e *Encoder) Value(v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		l := v.Len()
		for i := 0; i < l; i++ {
			e.Value(v.Index(i))
		}

	case reflect.Struct:
		t := v.Type()
		l := v.NumField()
		for i := 0; i < l; i++ {
			// see comment for corresponding code in decoder.value()
			if v := v.Field(i); v.CanSet() || t.Field(i).Name != "_" {
				e.Value(v)
			} else {
				e.Skip(v)
			}
		}

	case reflect.Slice:
		l := v.Len()
		for i := 0; i < l; i++ {
			e.Value(v.Index(i))
		}

	case reflect.Bool:
		e.Bool(v.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch v.Type().Kind() {
		case reflect.Int8:
			e.Int8(int8(v.Int()))
		case reflect.Int16:
			e.Int16(int16(v.Int()))
		case reflect.Int32:
			e.Int32(int32(v.Int()))
		case reflect.Int64:
			e.Int64(v.Int())
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		switch v.Type().Kind() {
		case reflect.Uint8:
			e.Uint8(uint8(v.Uint()))
		case reflect.Uint16:
			e.Uint16(uint16(v.Uint()))
		case reflect.Uint32:
			e.Uint32(uint32(v.Uint()))
		case reflect.Uint64:
			e.Uint64(v.Uint())
		}

	case reflect.Float32, reflect.Float64:
		switch v.Type().Kind() {
		case reflect.Float32:
			e.Uint32(math.Float32bits(float32(v.Float())))
		case reflect.Float64:
			e.Uint64(math.Float64bits(v.Float()))
		}

	case reflect.Complex64, reflect.Complex128:
		switch v.Type().Kind() {
		case reflect.Complex64:
			x := v.Complex()
			e.Uint32(math.Float32bits(float32(real(x))))
			e.Uint32(math.Float32bits(float32(imag(x))))
		case reflect.Complex128:
			x := v.Complex()
			e.Uint64(math.Float64bits(real(x)))
			e.Uint64(math.Float64bits(imag(x)))
		}
	}
}

func (d *Decoder) Skip(v reflect.Value) {
	d.offset += ValueSize(v)
}

func (e *Encoder) Skip(v reflect.Value) {
	n := ValueSize(v)
	zero := e.buf[e.offset : e.offset+n]
	for i := range zero {
		zero[i] = 0
	}
	e.offset += n
}

//

// Read reads structured binary data from r into data. Data must be a pointer to a fixed-size value or a slice of
// fixed-size values. Bytes read from r are decoded using the specified byte order and written to successive fields of
// the data. When decoding boolean values, a zero byte is decoded as false, and any other non-zero byte is decoded as
// true. When reading into structs, the field data for fields with blank (_) field names is skipped; i.e., blank field
// names may be used for padding. When reading into a struct, all non-blank fields must be exported or Read may panic.
//
// The error is EOF only if no bytes were read. If an EOF happens after reading some but not all the bytes, Read returns
// ErrUnexpectedEOF.
func Read(r io.Reader, order binary.ByteOrder, data any) error {
	// Fast path for basic types and slices.
	if n := IntDataSize(data); n != 0 {
		bs := make([]byte, n)
		if _, err := io.ReadFull(r, bs); err != nil {
			return err
		}
		switch data := data.(type) {
		case *bool:
			*data = bs[0] != 0
		case *int8:
			*data = int8(bs[0])
		case *uint8:
			*data = bs[0]
		case *int16:
			*data = int16(order.Uint16(bs))
		case *uint16:
			*data = order.Uint16(bs)
		case *int32:
			*data = int32(order.Uint32(bs))
		case *uint32:
			*data = order.Uint32(bs)
		case *int64:
			*data = int64(order.Uint64(bs))
		case *uint64:
			*data = order.Uint64(bs)
		case *float32:
			*data = math.Float32frombits(order.Uint32(bs))
		case *float64:
			*data = math.Float64frombits(order.Uint64(bs))
		case []bool:
			for i, x := range bs { // Easier to loop over the input for 8-bit values.
				data[i] = x != 0
			}
		case []int8:
			for i, x := range bs {
				data[i] = int8(x)
			}
		case []uint8:
			copy(data, bs)
		case []int16:
			for i := range data {
				data[i] = int16(order.Uint16(bs[2*i:]))
			}
		case []uint16:
			for i := range data {
				data[i] = order.Uint16(bs[2*i:])
			}
		case []int32:
			for i := range data {
				data[i] = int32(order.Uint32(bs[4*i:]))
			}
		case []uint32:
			for i := range data {
				data[i] = order.Uint32(bs[4*i:])
			}
		case []int64:
			for i := range data {
				data[i] = int64(order.Uint64(bs[8*i:]))
			}
		case []uint64:
			for i := range data {
				data[i] = order.Uint64(bs[8*i:])
			}
		case []float32:
			for i := range data {
				data[i] = math.Float32frombits(order.Uint32(bs[4*i:]))
			}
		case []float64:
			for i := range data {
				data[i] = math.Float64frombits(order.Uint64(bs[8*i:]))
			}
		default:
			n = 0 // fast path doesn't apply
		}
		if n != 0 {
			return nil
		}
	}

	// Fallback to reflect-based decoding.
	v := reflect.ValueOf(data)
	size := -1
	switch v.Kind() {
	case reflect.Pointer:
		v = v.Elem()
		size = ValueSize(v)
	case reflect.Slice:
		size = ValueSize(v)
	}
	if size < 0 {
		return errors.New("binary.Read: invalid type " + reflect.TypeOf(data).String())
	}
	d := &Decoder{order: order, buf: make([]byte, size)}
	if _, err := io.ReadFull(r, d.buf); err != nil {
		return err
	}
	d.Value(v)
	return nil
}

// Write writes the binary representation of data into w. Data must be a fixed-size value or a slice of fixed-size
// values, or a pointer to such data. Boolean values encode as one byte: 1 for true, and 0 for false. Bytes written to w
// are encoded using the specified byte order and read from successive fields of the data. When writing structs, zero
// values are written for fields with blank (_) field names.
func Write(w io.Writer, order binary.ByteOrder, data any) error {
	// Fast path for basic types and slices.
	if n := IntDataSize(data); n != 0 {
		bs := make([]byte, n)
		switch v := data.(type) {
		case *bool:
			if *v {
				bs[0] = 1
			} else {
				bs[0] = 0
			}
		case bool:
			if v {
				bs[0] = 1
			} else {
				bs[0] = 0
			}
		case []bool:
			for i, x := range v {
				if x {
					bs[i] = 1
				} else {
					bs[i] = 0
				}
			}
		case *int8:
			bs[0] = byte(*v)
		case int8:
			bs[0] = byte(v)
		case []int8:
			for i, x := range v {
				bs[i] = byte(x)
			}
		case *uint8:
			bs[0] = *v
		case uint8:
			bs[0] = v
		case []uint8:
			bs = v
		case *int16:
			order.PutUint16(bs, uint16(*v))
		case int16:
			order.PutUint16(bs, uint16(v))
		case []int16:
			for i, x := range v {
				order.PutUint16(bs[2*i:], uint16(x))
			}
		case *uint16:
			order.PutUint16(bs, *v)
		case uint16:
			order.PutUint16(bs, v)
		case []uint16:
			for i, x := range v {
				order.PutUint16(bs[2*i:], x)
			}
		case *int32:
			order.PutUint32(bs, uint32(*v))
		case int32:
			order.PutUint32(bs, uint32(v))
		case []int32:
			for i, x := range v {
				order.PutUint32(bs[4*i:], uint32(x))
			}
		case *uint32:
			order.PutUint32(bs, *v)
		case uint32:
			order.PutUint32(bs, v)
		case []uint32:
			for i, x := range v {
				order.PutUint32(bs[4*i:], x)
			}
		case *int64:
			order.PutUint64(bs, uint64(*v))
		case int64:
			order.PutUint64(bs, uint64(v))
		case []int64:
			for i, x := range v {
				order.PutUint64(bs[8*i:], uint64(x))
			}
		case *uint64:
			order.PutUint64(bs, *v)
		case uint64:
			order.PutUint64(bs, v)
		case []uint64:
			for i, x := range v {
				order.PutUint64(bs[8*i:], x)
			}
		case *float32:
			order.PutUint32(bs, math.Float32bits(*v))
		case float32:
			order.PutUint32(bs, math.Float32bits(v))
		case []float32:
			for i, x := range v {
				order.PutUint32(bs[4*i:], math.Float32bits(x))
			}
		case *float64:
			order.PutUint64(bs, math.Float64bits(*v))
		case float64:
			order.PutUint64(bs, math.Float64bits(v))
		case []float64:
			for i, x := range v {
				order.PutUint64(bs[8*i:], math.Float64bits(x))
			}
		}
		_, err := w.Write(bs)
		return err
	}

	// Fallback to reflect-based encoding.
	v := reflect.Indirect(reflect.ValueOf(data))
	size := ValueSize(v)
	if size < 0 {
		return errors.New("binary.Write: invalid type " + reflect.TypeOf(data).String())
	}
	buf := make([]byte, size)
	e := &Encoder{order: order, buf: buf}
	e.Value(v)
	_, err := w.Write(buf)
	return err
}
