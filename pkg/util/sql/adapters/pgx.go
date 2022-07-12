package adapters

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	sqb "github.com/wrmsr/bane/pkg/util/sql/base"
)

//

type PgxRows struct {
	sqb.BaseRows

	r  pgx.Rows
	cs []sqb.Column
}

func newPgxRows(r pgx.Rows) *PgxRows {
	return &PgxRows{
		r: r,
	}
}

var _ sqb.Rows = &PgxRows{}

func (r *PgxRows) Columns() ([]sqb.Column, error) {
	if r.cs == nil {
		fis := r.r.FieldDescriptions()
		cs := make([]sqb.Column, len(fis))
		for i, fi := range fis {
			// FIXME:
			//fi.DataTypeOID
			cs[i] = sqb.Column{
				Name: string(fi.Name),
				//DbType: ct.DatabaseTypeName(),
				//Type:   ct.ScanType(),
			}
		}
		r.cs = cs
	}
	return r.cs, nil
}

func (r PgxRows) Next() bool {
	return r.r.Next()
}

func (r PgxRows) Scan(dest ...any) error {
	return r.r.Scan(dest...)
}

func (r PgxRows) Err() error {
	return r.r.Err()
}

func (r PgxRows) Close() error {
	r.r.Close()
	return nil
}

//

type PgxConn struct {
	sqb.BaseConn

	c *pgxpool.Conn
}

var _ sqb.Conn = PgxConn{}

func (c PgxConn) Adapter() sqb.Adapter { return PgxAdapter{} }

func (c PgxConn) Query(qry sqb.Query) (sqb.Rows, error) {
	switch qry.Mode {
	case sqb.QueryQuery:
		r, err := c.c.Query(qry.Ctx, qry.Text, qry.Args...)
		if err != nil {
			return nil, err
		}
		return newPgxRows(r), nil
	case sqb.ExecQuery:
		_, err := c.c.Exec(qry.Ctx, qry.Text, qry.Args...)
		return nil, err
	}
	panic("unreachable")
}

func (c PgxConn) Close() error {
	c.c.Release()
	return nil
}

//

type PgxDb struct {
	sqb.BaseDb

	d *pgxpool.Pool
}

var _ sqb.Db = PgxDb{}

func NewPgxDb(d *pgxpool.Pool) PgxDb {
	return PgxDb{d: d}
}

func (d PgxDb) Adapter() sqb.Adapter { return PgxAdapter{} }

func (d PgxDb) Connect(ctx context.Context) (sqb.Conn, error) {
	c, err := d.d.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	return PgxConn{c: c}, nil
}

func (d PgxDb) Query(qry sqb.Query) (sqb.Rows, error) {
	switch qry.Mode {
	case sqb.QueryQuery:
		r, err := d.d.Query(qry.Ctx, qry.Text, qry.Args...)
		if err != nil {
			return nil, err
		}
		return newPgxRows(r), nil
	case sqb.ExecQuery:
		_, err := d.d.Exec(qry.Ctx, qry.Text, qry.Args...)
		return nil, err
	}
	panic("unreachable")
}

func (d PgxDb) Close() error {
	d.d.Close()
	return nil
}

//

type PgxAdapter struct {
	sqb.BaseAdapter
}

var _ sqb.Adapter = PgxAdapter{}
