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

func (t Type) String() string {
	switch t {
	case InvalidType:
		return "invalid"
	case IntegerType:
		return "integer"
	case IdentifierType:
		return "identifier"
	case DateType:
		return "date"
	case DoubleType:
		return "double"
	case VarcharType:
		return "varchar"
	}
	panic("unreachable")
}

//

type Column struct {
	name string
	ty   Type

	precision int
	scale     int
}

func (c Column) Type() Type   { return c.ty }
func (c Column) Name() string { return c.name }

func (c Column) Precision() int { return c.precision }
func (c Column) Scale() int     { return c.scale }

//

type Entity struct {
	name string

	primaryKey string

	C       ctr.OrdMap[string, Column]
	indexes ctr.Set[string]
}

func NewEntity(
	name string,
	primaryKey string,
	columns []Column,
	indexes []string,
) *Entity {
	return &Entity{
		name: check.NotZero(name),

		primaryKey: check.NotZero(primaryKey),

		C:       ctr.NewOrdMap(its.Map(its.OfSlice(columns), bt.KvMaker(Column.Name))),
		indexes: ctr.NewSet(its.OfSlice(indexes)),
	}
}

func (e Entity) Name() string { return e.name }

func (e Entity) PrimaryKey() string { return e.primaryKey }

func (e Entity) Columns() ctr.OrdMap[string, Column] { return e.C }
func (e Entity) Indexes() ctr.Set[string]            { return e.indexes }

//

var (
	Region = NewEntity(
		"region",
		"region_key",
		[]Column{
			{name: "r_regionkey", ty: IdentifierType},
			{name: "r_name", ty: VarcharType, precision: 25},
			{name: "r_comment", ty: VarcharType, precision: 152},
		},
		nil,
	)
)
