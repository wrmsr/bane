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

	primaryKey ctr.List[string]
	indexes    ctr.List[ctr.List[string]]

	columns ctr.OrdMap[string, Column]
}

func NewEntity(
	name string,
	primaryKey []string,
	indexes [][]string,
	columns []Column,
) *Entity {
	e := &Entity{
		name: check.NotZero(name),

		primaryKey: ctr.NewList(its.OfSlice(primaryKey)),
		indexes: ctr.NewList(its.Map(its.OfSlice(indexes), func(s []string) ctr.List[string] {
			return ctr.NewList(its.OfSlice(s))
		})),

		columns: ctr.NewOrdMap(its.CheckUniqueKeys(its.Map(its.OfSlice(columns), bt.KvMaker(Column.Name)))),
	}

	its.ForAll(its.AsTraversable(its.Values[string, Column](e.columns)), func(c Column) {
		check.NotZero(c.name)
	})

	its.ForAll[string](e.primaryKey, func(cn string) {
		check.Contains[string](e.columns, cn)
	})

	its.ForAll[ctr.List[string]](e.indexes, func(idx ctr.List[string]) {
		check.NotEmpty(idx)
		its.ForAll[string](idx, func(cn string) {
			check.Contains[string](e.columns, cn)
		})
	})

	return e
}

func (e Entity) Name() string { return e.name }

func (e Entity) PrimaryKey() ctr.List[string]        { return e.primaryKey }
func (e Entity) Indexes() ctr.List[ctr.List[string]] { return e.indexes }

func (e Entity) Columns() ctr.OrdMap[string, Column] { return e.columns }

//

type EntitySet struct {
	entities ctr.Map[string, *Entity]
}

func NewEntitySet(entities []*Entity) *EntitySet {
	return &EntitySet{
		entities: ctr.NewMap(its.Map(its.OfSlice(entities), bt.KvMaker((*Entity).Name))),
	}
}

func (es EntitySet) Entities() ctr.Map[string, *Entity] { return es.entities }

//

var (
	Entities = NewEntitySet([]*Entity{
		NewEntity(
			"region",
			[]string{"r_regionkey"},
			nil,
			[]Column{
				{name: "r_regionkey", ty: IdentifierType},
				{name: "r_name", ty: VarcharType, precision: 25},
				{name: "r_comment", ty: VarcharType, precision: 152},
			},
		),

		NewEntity(
			"nation",
			[]string{"n_nationkey"},
			[][]string{{"n_regionkey"}},
			[]Column{
				{name: "n_nationkey", ty: IdentifierType},
				{name: "n_name", ty: VarcharType, precision: 25},
				{name: "n_regionkey", ty: IdentifierType},
				{name: "n_comment", ty: VarcharType, precision: 152},
			},
		),

		NewEntity(
			"part",
			[]string{"p_partkey"},
			nil,
			[]Column{
				{name: "p_partkey", ty: IdentifierType},
				{name: "p_name", ty: VarcharType, precision: 55},
				{name: "p_mfgr", ty: VarcharType, precision: 25},
				{name: "p_brand", ty: VarcharType, precision: 10},
				{name: "p_type", ty: VarcharType, precision: 25},
				{name: "p_size", ty: IntegerType},
				{name: "p_container", ty: VarcharType, precision: 10},
				{name: "p_retailprice", ty: DoubleType},
				{name: "p_comment", ty: VarcharType, precision: 23},
			},
		),

		NewEntity(
			"supplier",
			[]string{"s_suppkey"},
			[][]string{{"s_nationkey"}},
			[]Column{
				{name: "s_suppkey", ty: IdentifierType},
				{name: "s_name", ty: VarcharType, precision: 25},
				{name: "s_address", ty: VarcharType, precision: 40},
				{name: "s_nationkey", ty: IdentifierType},
				{name: "s_phone", ty: VarcharType, precision: 15},
				{name: "s_acctbal", ty: DoubleType},
				{name: "s_comment", ty: VarcharType, precision: 101},
			},
		),

		NewEntity(
			"part_supplier",
			[]string{"ps_suppkey", "ps_partkey"},
			nil,
			[]Column{
				{name: "ps_partkey", ty: IdentifierType},
				{name: "ps_suppkey", ty: IdentifierType},
				{name: "ps_availqty", ty: IntegerType},
				{name: "ps_supplycost", ty: DoubleType},
				{name: "ps_comment", ty: VarcharType, precision: 199},
			},
		),

		NewEntity(
			"customer",
			[]string{"c_custkey"},
			[][]string{{"c_nationkey"}},
			[]Column{
				{name: "c_custkey", ty: IdentifierType},
				{name: "c_name", ty: VarcharType, precision: 25},
				{name: "c_address", ty: VarcharType, precision: 40},
				{name: "c_nationkey", ty: IdentifierType},
				{name: "c_phone", ty: VarcharType, precision: 25},
				{name: "c_acctbal", ty: DoubleType},
				{name: "c_mktsegment", ty: VarcharType, precision: 10},
				{name: "c_comment", ty: VarcharType, precision: 117},
			},
		),

		NewEntity(
			"order",
			[]string{"o_orderkey"},
			[][]string{{"o_custkey"}},
			[]Column{
				{name: "o_orderkey", ty: IdentifierType},
				{name: "o_custkey", ty: IdentifierType},
				{name: "o_orderstatus", ty: VarcharType, precision: 1},
				{name: "o_totalprice", ty: DoubleType},
				{name: "o_orderdate", ty: DateType},
				{name: "o_orderpriority", ty: VarcharType, precision: 15},
				{name: "o_clerk", ty: VarcharType, precision: 15},
				{name: "o_shippriority", ty: IntegerType},
				{name: "o_comment", ty: VarcharType, precision: 79},
			},
		),

		NewEntity(
			"line_item",
			[]string{"l_orderkey", "l_linenumber"},
			[][]string{{"l_partkey"}, {"l_suppkey"}},
			[]Column{
				{name: "l_orderkey", ty: IdentifierType},
				{name: "l_partkey", ty: IdentifierType},
				{name: "l_suppkey", ty: IdentifierType},
				{name: "l_linenumber", ty: IntegerType},
				{name: "l_quantity", ty: DoubleType},
				{name: "l_extendedprice", ty: DoubleType},
				{name: "l_discount", ty: DoubleType},
				{name: "l_tax", ty: DoubleType},
				{name: "l_returnflag", ty: VarcharType, precision: 1},
				{name: "l_linestatus", ty: VarcharType, precision: 1},
				{name: "l_shipdate", ty: DateType},
				{name: "l_commitdate", ty: DateType},
				{name: "l_receiptdate", ty: DateType},
				{name: "l_shipinstruct", ty: VarcharType, precision: 25},
				{name: "l_shipmode", ty: VarcharType, precision: 10},
				{name: "l_comment", ty: VarcharType, precision: 44},
			},
		),
	})
)
