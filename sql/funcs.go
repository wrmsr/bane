package sql

import (
	"context"

	eu "github.com/wrmsr/bane/core/errors"
	bt "github.com/wrmsr/bane/core/types"
	sqb "github.com/wrmsr/bane/sql/base"
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

func Iter(ctx context.Context, o sqb.Querier, objs ...any) (*RowsIterator, error) {
	r, err := Query(ctx, o, objs...)
	if err != nil {
		return nil, err
	}
	return IterRows(r), nil
}

func All(ctx context.Context, o sqb.Querier, objs ...any) (s []Row, err error) {
	ri, err := Iter(ctx, o, objs...)
	if err != nil {
		return nil, err
	}
	defer eu.AppendInvoke(&err, eu.Close(ri))

	s = make([]Row, 0, 16)
	for ri.HasNext() {
		rr := ri.Next()
		if rr.Err != nil {
			s = nil
			err = rr.Err
			return
		}
		s = append(s, rr.Val.Clone())
	}
	return
}

func maybe(ctx context.Context, o sqb.Querier, objs []any, onlyOne bool) (bt.Optional[Row], error) {
	ri, err := Iter(ctx, o, objs...)
	if err != nil {
		return bt.None[Row](), err
	}
	defer eu.AppendInvoke(&err, eu.Close(ri))

	if !ri.HasNext() {
		return bt.None[Row](), nil
	}

	re := ri.Next()
	if re.Err != nil {
		return bt.None[Row](), re.Err
	}

	if onlyOne && ri.HasNext() {
		return bt.None[Row](), MultipleFoundError{}
	}

	return bt.Just(re.Val), nil
}

func MaybeFirst(ctx context.Context, o sqb.Querier, objs ...any) (bt.Optional[Row], error) {
	return maybe(ctx, o, objs, false)
}

func MaybeOne(ctx context.Context, o sqb.Querier, objs ...any) (bt.Optional[Row], error) {
	return maybe(ctx, o, objs, true)
}

func unwrapMaybe(ro bt.Optional[Row], err error) (Row, error) {
	if err != nil {
		return Row{}, err
	}
	if !ro.Present() {
		return Row{}, NotFoundError{}
	}
	return ro.Value(), nil
}

func First(ctx context.Context, o sqb.Querier, objs ...any) (Row, error) {
	return unwrapMaybe(maybe(ctx, o, objs, false))
}

func One(ctx context.Context, o sqb.Querier, objs ...any) (Row, error) {
	return unwrapMaybe(maybe(ctx, o, objs, true))
}

func Scalar(ctx context.Context, o sqb.Querier, objs ...any) (any, error) {
	r, err := One(ctx, o, objs...)
	if err != nil {
		return nil, err
	}

	if r.Len() < 1 {
		return nil, NotEnoughValuesError{}
	}
	if r.Len() > 1 {
		return nil, TooManyValuesError{}
	}
	return r.Value(0), nil
}
