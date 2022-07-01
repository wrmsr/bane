package io

type StringWriter interface {
	WriteString(string) (int, error)
}
