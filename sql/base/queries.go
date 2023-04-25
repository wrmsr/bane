package base

import "context"

type QueryMode int8

const (
	QueryQuery QueryMode = iota
	ExecQuery
)

type Query struct {
	Ctx  context.Context
	Mode QueryMode
	Text string
	Args []any
}
