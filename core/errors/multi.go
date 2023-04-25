// Copyright (c) 2017-2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package errors

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"sync/atomic"
)

var (
	singleLineSep = []byte("; ")

	multiLinePrefix = []byte("the following errors occurred:")
	multiLineSep    = []byte("\n -  ")
	multiLineIndent = []byte("    ")
)

var _bufferPool = sync.Pool{
	New: func() any {
		return &bytes.Buffer{}
	},
}

type errorGroup interface {
	Errors() []error
}

func Errors(err error) []error {
	if err == nil {
		return nil
	}

	eg, ok := err.(*multiError)
	if !ok {
		return []error{err}
	}

	return append(([]error)(nil), eg.Errors()...)
}

type multiError struct {
	copyNeeded uint32
	errors     []error
}

var _ errorGroup = (*multiError)(nil)

func (me *multiError) Errors() []error {
	if me == nil {
		return nil
	}
	return me.errors
}

func (me *multiError) As(target any) bool {
	for _, err := range me.Errors() {
		if errors.As(err, target) {
			return true
		}
	}
	return false
}

func (me *multiError) Is(target error) bool {
	for _, err := range me.Errors() {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

func (me *multiError) Error() string {
	if me == nil {
		return ""
	}

	buff := _bufferPool.Get().(*bytes.Buffer)
	buff.Reset()

	me.writeSingleline(buff)

	result := buff.String()
	_bufferPool.Put(buff)
	return result
}

func (me *multiError) Format(f fmt.State, c rune) {
	if c == 'v' && f.Flag('+') {
		me.writeMultiline(f)
	} else {
		me.writeSingleline(f)
	}
}

func (me *multiError) writeSingleline(w io.Writer) {
	first := true
	for _, item := range me.errors {
		if first {
			first = false
		} else {
			_, _ = w.Write(singleLineSep)
		}
		_, _ = io.WriteString(w, item.Error())
	}
}

func (me *multiError) writeMultiline(w io.Writer) {
	_, _ = w.Write(multiLinePrefix)
	for _, item := range me.errors {
		_, _ = w.Write(multiLineSep)
		writePrefixLine(w, multiLineIndent, fmt.Sprintf("%+v", item))
	}
}

func writePrefixLine(w io.Writer, prefix []byte, s string) {
	first := true
	for len(s) > 0 {
		if first {
			first = false
		} else {
			_, _ = w.Write(prefix)
		}

		idx := strings.IndexByte(s, '\n')
		if idx < 0 {
			idx = len(s) - 1
		}

		_, _ = io.WriteString(w, s[:idx+1])
		s = s[idx+1:]
	}
}

type inspectResult struct {
	Count              int
	Capacity           int
	FirstErrorIdx      int
	ContainsMultiError bool
}

func inspect(errors []error) (res inspectResult) {
	first := true
	for i, err := range errors {
		if err == nil {
			continue
		}

		res.Count++
		if first {
			first = false
			res.FirstErrorIdx = i
		}

		if merr, ok := err.(*multiError); ok {
			res.Capacity += len(merr.errors)
			res.ContainsMultiError = true
		} else {
			res.Capacity++
		}
	}
	return
}

func fromSlice(errors []error) error {
	switch len(errors) {
	case 0:
		return nil
	case 1:
		return errors[0]
	}

	res := inspect(errors)
	switch res.Count {
	case 0:
		return nil
	case 1:
		return errors[res.FirstErrorIdx]
	case len(errors):
		if !res.ContainsMultiError {
			out := append(([]error)(nil), errors...)
			return &multiError{errors: out}
		}
	}

	nonNilErrs := make([]error, 0, res.Capacity)
	for _, err := range errors[res.FirstErrorIdx:] {
		if err == nil {
			continue
		}

		if nested, ok := err.(*multiError); ok {
			nonNilErrs = append(nonNilErrs, nested.errors...)
		} else {
			nonNilErrs = append(nonNilErrs, err)
		}
	}

	return &multiError{errors: nonNilErrs}
}

func Combine(errors ...error) error {
	return fromSlice(errors)
}

func Append(left error, right error) error {
	switch {
	case left == nil:
		return right
	case right == nil:
		return left
	}

	if _, ok := right.(*multiError); !ok {
		if l, ok := left.(*multiError); ok && atomic.SwapUint32(&l.copyNeeded, 1) == 0 {
			errs := append(l.errors, right)
			return &multiError{errors: errs}
		} else if !ok {
			// Both errors are single errors.
			return &multiError{errors: []error{left, right}}
		}
	}

	es := [2]error{left, right}
	return fromSlice(es[0:])
}

func AppendInto(into *error, err error) (errored bool) {
	if err == nil {
		return false
	}
	*into = Append(*into, err)
	return true
}

type Invoker interface {
	Invoke() error
}

type Invoke func() error

func (i Invoke) Invoke() error { return i() }

func Close(closer io.Closer) Invoker {
	return Invoke(closer.Close)
}

func AppendInvoke(into *error, invoker Invoker) {
	AppendInto(into, invoker.Invoke())
}
