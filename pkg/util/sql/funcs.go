package sql

import (
	"context"

	sqb "github.com/wrmsr/bane/pkg/util/sql/base"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func MakeQuery(ctx context.Context, mode sqb.QueryMode, ad sqb.Adapter, objs ...any) (sqb.Query, error) {
	if len(objs) == 1 && bt.Is[string](objs[0]) {
		return sqb.Query{
			Ctx:  ctx,
			Mode: mode,
			Text: objs[0].(string),
		}, nil
	}
	panic("fixme")
}

func Query(ctx context.Context, o sqb.Querier, objs ...any) (sqb.Rows, error) {
	q, err := MakeQuery(ctx, sqb.QueryQuery, o.Adapter(), objs...)
	if err != nil {
		return nil, err
	}
	return o.Query(q)
}

func Exec(ctx context.Context, o sqb.Querier, objs ...any) error {
	q, err := MakeQuery(ctx, sqb.ExecQuery, o.Adapter(), objs...)
	if err != nil {
		return err
	}
	_, err = o.Query(q)
	return err
}

func All(ctx context.Context, o sqb.Querier, objs ...any) ([]Row, error) {
	panic("implement me")
}

func First(ctx context.Context, o sqb.Querier, objs ...any) (Row, error) {
	panic("implement me")
}

func One(ctx context.Context, o sqb.Querier, objs ...any) (Row, error) {
	panic("implement me")
}

func Scalar(ctx context.Context, o sqb.Querier, objs ...any) (any, error) {
	panic("implement me")
}
