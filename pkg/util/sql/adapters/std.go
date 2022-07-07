package adapters

import (
	"context"
	"database/sql"

	sqb "github.com/wrmsr/bane/pkg/util/sql/base"
)

//

type StdRows struct {
	sqb.BaseRows

	r  *sql.Rows
	cs []sqb.Column
}

func newStdRows(r *sql.Rows) *StdRows {
	return &StdRows{
		r: r,
	}
}

var _ sqb.Rows = &StdRows{}

func (r *StdRows) Columns() ([]sqb.Column, error) {
	if r.cs == nil {
		cns, err := r.r.Columns()
		if err != nil {
			return nil, err
		}
		cts, err := r.r.ColumnTypes()
		if err != nil {
			return nil, err
		}
		cs := make([]sqb.Column, len(cns))
		for i, cn := range cns {
			ct := cts[i]
			cs[i] = sqb.Column{
				Name:   cn,
				DbType: ct.DatabaseTypeName(),
				Type:   ct.ScanType(),
			}
		}
		r.cs = cs
	}
	return r.cs, nil
}

func (r StdRows) Next() bool {
	return r.r.Next()
}

func (r StdRows) Scan(dest ...any) error {
	return r.r.Scan(dest...)
}

func (r StdRows) Err() error {
	return r.r.Err()
}

func (r StdRows) Close() error {
	return r.r.Close()
}

//

type StdConn struct {
	sqb.BaseConn

	c *sql.Conn
}

var _ sqb.Conn = StdConn{}

func (s StdConn) Adapter() sqb.Adapter { return StdAdapter{} }

func (s StdConn) Query(qry sqb.Query) (sqb.Rows, error) {
	switch qry.Mode {
	case sqb.QueryQuery:
		r, err := s.c.QueryContext(qry.Ctx, qry.Text, qry.Args...)
		if err != nil {
			return nil, err
		}
		return newStdRows(r), nil
	case sqb.ExecQuery:
		_, err := s.c.ExecContext(qry.Ctx, qry.Text, qry.Args...)
		return nil, err
	}
	panic("unreachable")
}

func (s StdConn) Close() error {
	return s.c.Close()
}

//

type StdDb struct {
	sqb.BaseDb

	d *sql.DB
}

var _ sqb.Db = StdDb{}

func NewStdDb(d *sql.DB) StdDb {
	return StdDb{d: d}
}

func (d StdDb) Adapter() sqb.Adapter { return StdAdapter{} }

func (d StdDb) Connect(ctx context.Context) (sqb.Conn, error) {
	c, err := d.d.Conn(ctx)
	if err != nil {
		return nil, err
	}
	return StdConn{c: c}, nil
}

func (d StdDb) Query(qry sqb.Query) (sqb.Rows, error) {
	switch qry.Mode {
	case sqb.QueryQuery:
		r, err := d.d.QueryContext(qry.Ctx, qry.Text, qry.Args...)
		if err != nil {
			return nil, err
		}
		return newStdRows(r), nil
	case sqb.ExecQuery:
		_, err := d.d.ExecContext(qry.Ctx, qry.Text, qry.Args...)
		return nil, err
	}
	panic("unreachable")
}

func (d StdDb) Close() error {
	return d.d.Close()
}

//

type StdAdapter struct {
	sqb.BaseAdapter
}

var _ sqb.Adapter = StdAdapter{}
