package tpch

import (
	"github.com/wrmsr/bane/pkg/util/check"
	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Type int8

const (
	InvalidType Type = iota
	IntegerType
	IdentifierType
	DateType
	DoubleType
	VarcharType
)

//

type Column struct {
	name string
	ty   Type

	precision int
	scale     int
}

func (c Column) Ty() Type     { return c.ty }
func (c Column) Name() string { return c.name }

func (c Column) Precision() int { return c.precision }
func (c Column) Scale() int     { return c.scale }

//

type Entity struct {
	Name    string
	Columns ctr.OrdMap[string, Column]
}

func NewEntity(name string, columns []Column) *Entity {
	return &Entity{
		Name:    check.NotZero(name),
		Columns: ctr.NewOrdMap(its.Map(its.OfSlice(columns), bt.KvMaker(Column.Name))),
	}
}
