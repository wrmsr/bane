package adapters

import (
	"database/sql"

	bsb "github.com/wrmsr/bane/pkg/utils/sql/base"
)

//

type StdRows struct {
	bsb.BaseRows

	r  *sql.Rows
	cs []bsb.Column
}

func newStdRows(r *sql.Rows) *StdRows {
	return &StdRows{
		r: r,
	}
}

var _ bsb.Rows = &StdRows{}

func (r *StdRows) Columns() ([]bsb.Column, error) {
	if r.cs == nil {
		cns, err := r.r.Columns()
		if err != nil {
			return nil, err
		}
		cts, err := r.r.ColumnTypes()
		if err != nil {
			return nil, err
		}
		cs := make([]bsb.Column, len(cns))
		for i, cn := range cns {
			ct := cts[i]
			cs[i] = bsb.Column{
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
	bsb.BaseConn

	c *sql.Conn
}

var _ bsb.Conn = StdConn{}

func (s StdConn) Adapter() bsb.Adapter { return StdAdapter{} }

func (s StdConn) Query(qry bsb.Query) (bsb.Rows, error) {
	switch qry.Mode {
	case bsb.QueryQuery:
		r, err := s.c.QueryContext(qry.Ctx, qry.Text, qry.Args...)
		if err != nil {
			return nil, err
		}
		return newStdRows(r), nil
	case bsb.ExecQuery:
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
	bsb.BaseDb

	d *sql.DB
}

var _ bsb.Db = StdDb{}

func (d StdDb) Adapter() bsb.Adapter { return StdAdapter{} }

func (d StdDb) Query(qry bsb.Query) (bsb.Rows, error) {
	switch qry.Mode {
	case bsb.QueryQuery:
		r, err := d.d.QueryContext(qry.Ctx, qry.Text, qry.Args...)
		if err != nil {
			return nil, err
		}
		return newStdRows(r), nil
	case bsb.ExecQuery:
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
	bsb.BaseAdapter
}

var _ bsb.Adapter = StdAdapter{}
