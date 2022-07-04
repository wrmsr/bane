package sql

import (
	"context"

	sqb "github.com/wrmsr/bane/pkg/util/sql/base"
)

func MakeQuery(ctx context.Context, ad sqb.Adapter, objs ...any) (sqb.Query, error) {
	panic("implement me")
}

func Query(ctx context.Context, b sqb.Querier, objs ...any) {
	panic("implement me")
}

func Exec(ctx context.Context, o sqb.Querier, objs ...any) error {
	panic("implement me")
}

func All() ([]Row, error) {
	panic("implement me")
}

func First() (Row, error) {
	panic("implement me")
}

func One() (Row, error) {
	panic("implement me")
}

func Scalar() (any, error) {
	panic("implement me")
}
