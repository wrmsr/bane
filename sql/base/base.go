package base

import (
	"context"
	"reflect"
)

//

type Querier interface {
	Adapter() Adapter
	Query(qry Query) (Rows, error)

	Close() error
}

//

type Rows interface {
	Columns() (*Columns, error)

	Next() bool
	Scan(dest ...any) error
	Err() error

	Close() error

	isRows()
}

type BaseRows struct{}

func (r BaseRows) isRows() {}

//

type Conn interface {
	Querier

	isConn()
}

type BaseConn struct{}

func (c BaseConn) isConn() {}

//

type Db interface {
	Querier

	Connect(ctx context.Context) (Conn, error)

	isDb()
}

type BaseDb struct{}

func (d BaseDb) isDb() {}

//

type Adapter interface {
	isAdapter()

	ScanType(c Column) reflect.Type
}

type BaseAdapter struct{}

func (a BaseAdapter) isAdapter() {}

func (a BaseAdapter) ScanType(c Column) reflect.Type {
	panic("implement me")
}
