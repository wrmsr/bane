package tpch

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	au "github.com/wrmsr/bane/pkg/util/atomic"
	"github.com/wrmsr/bane/pkg/util/check"
	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Identifier int64

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

	fld reflect.StructField
}

func (c Column) Type() Type   { return c.ty }
func (c Column) Name() string { return c.name }

func (c Column) Precision() int { return c.precision }
func (c Column) Scale() int     { return c.scale }

func reflectColumn(fld reflect.StructField) Column {
	col := Column{
		fld: fld,
	}

	switch fld.Type {
	case rfl.Type[int64]():
		col.ty = IntegerType
	case rfl.Type[Identifier]():
		col.ty = IdentifierType
	case rfl.Type[time.Time]():
		col.ty = DateType
	case rfl.Type[float64]():
		col.ty = DoubleType
	case rfl.Type[string]():
		col.ty = VarcharType
	default:
		panic(fmt.Errorf("unhandled entity field type: %v", fld.Type))
	}

	tag := fld.Tag.Get("tpch")
	if strings.Contains(tag, ",") {
		var prec string
		tag, prec, _ = strings.Cut(tag, ",")
		col.precision = check.Must1(strconv.Atoi(prec))
	}

	col.name = tag
	return col
}

//

type Entity struct {
	name string

	primaryKey ctr.List[string]
	indexes    ctr.List[ctr.List[string]]

	columns ctr.OrderedMap[string, Column]

	ty reflect.Type
}

func reflectEntity(
	ty reflect.Type,
	name string,
	primaryKey []string,
	indexes [][]string,
) *Entity {
	e := &Entity{
		name: check.NotZero(name),

		primaryKey: ctr.NewList(its.OfSlice(primaryKey)),
		indexes: ctr.NewList(its.Map(its.OfSlice(indexes), func(s []string) ctr.List[string] {
			return ctr.NewList(its.OfSlice(s))
		})),

		columns: ctr.NewOrderedMap(its.Map(
			its.Range(0, ty.NumField(), 1),
			func(i int) bt.Kv[string, Column] {
				col := reflectColumn(ty.Field(i))
				return bt.KvOf(col.name, col)
			}),
		),

		ty: ty,
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

func (e Entity) Columns() ctr.OrderedMap[string, Column] { return e.columns }

//

type EntitySet struct {
	entities ctr.Map[string, *Entity]
}

func newEntitySet(entities []*Entity) *EntitySet {
	return &EntitySet{
		entities: ctr.NewMap(its.CheckUniqueKeys(its.Map(its.OfSlice(entities), bt.KvMaker((*Entity).Name)))),
	}
}

func (es EntitySet) Entities() ctr.Map[string, *Entity] { return es.entities }

//

var _entities []*Entity

func _registerEntity(e *Entity) any {
	_entities = append(_entities, e)
	return _entities
}

var entitiesOnce = au.NewLazy(func() *EntitySet {
	return newEntitySet(_entities)
})

func getEntities() *EntitySet {
	return entitiesOnce.Get()
}

//

var _ = _registerEntity(reflectEntity(
	rfl.Type[Region](),
	"region",
	[]string{"r_regionkey"},
	nil,
))

type Region struct {
	RegionKey Identifier `tpch:"r_regionkey"`
	Name      string     `tpch:"r_name,25"`
	Comment   string     `tpch:"r_comment,152"`
}

//

var _ = _registerEntity(reflectEntity(
	rfl.Type[Nation](),
	"nation",
	[]string{"n_nationkey"},
	[][]string{{"n_regionkey"}},
))

type Nation struct {
	NationKey Identifier `tpch:"n_nationkey"`
	Name      string     `tpch:"n_name,25"`
	RegionKey Identifier `tpch:"n_regionkey"`
	Comment   string     `tpch:"n_comment,152"`
}

//

var _ = _registerEntity(reflectEntity(
	rfl.Type[Part](),
	"part",
	[]string{"p_partkey"},
	nil,
))

type Part struct {
	PartKey      Identifier `tpch:"p_partkey"`
	Name         string     `tpch:"p_name,55"`
	Manufacturer string     `tpch:"p_mfgr,25"`
	Brand        string     `tpch:"p_brand,10"`
	Type         string     `tpch:"p_type,25"`
	Size         int64      `tpch:"p_size"`
	Container    string     `tpch:"p_container,10"`
	RetailPrice  float64    `tpch:"p_retailprice"`
	Comment      string     `tpch:"p_comment,23"`
}

//

var _ = _registerEntity(reflectEntity(
	rfl.Type[Supplier](),
	"supplier",
	[]string{"s_suppkey"},
	[][]string{{"s_nationkey"}},
))

type Supplier struct {
	SupplierKey    Identifier `tpch:"s_suppkey"`
	Name           string     `tpch:"s_name,25"`
	Address        string     `tpch:"s_address,40"`
	NationKey      Identifier `tpch:"s_nationkey"`
	Phone          string     `tpch:"s_phone,15"`
	AccountBalance float64    `tpch:"s_acctbal"`
	Comment        string     `tpch:"s_comment,101"`
}

//

var _ = _registerEntity(reflectEntity(
	rfl.Type[PartSupplier](),
	"part_supplier",
	[]string{"ps_suppkey", "ps_partkey"},
	nil,
))

type PartSupplier struct {
	PartKey           Identifier `tpch:"ps_partkey"`
	SupplierKey       Identifier `tpch:"ps_suppkey"`
	AvailableQuantity int64      `tpch:"ps_availqty"`
	SupplyCost        float64    `tpch:"ps_supplycost"`
	Comment           string     `tpch:"ps_comment,199"`
}

//

var _ = _registerEntity(reflectEntity(
	rfl.Type[Customer](),
	"customer",
	[]string{"c_custkey"},
	[][]string{{"c_nationkey"}},
))

type Customer struct {
	CustomerKey    Identifier `tpch:"c_custkey"`
	Name           string     `tpch:"c_name,25"`
	Address        string     `tpch:"c_address,40"`
	NationKey      Identifier `tpch:"c_nationkey"`
	Phone          string     `tpch:"c_phone,25"`
	AccountBalance float64    `tpch:"c_acctbal"`
	MarketSegment  string     `tpch:"c_mktsegment,10"`
	Comment        string     `tpch:"c_comment,117"`
}

//

var _ = _registerEntity(reflectEntity(
	rfl.Type[Order](),
	"order",
	[]string{"o_orderkey"},
	[][]string{{"o_custkey"}},
))

type Order struct {
	OrderKey      Identifier `tpch:"o_orderkey"`
	CustomerKey   Identifier `tpch:"o_custkey"`
	OrderStatus   string     `tpch:"o_orderstatus,1"`
	TotalPrice    float64    `tpch:"o_totalprice"`
	OrderDate     time.Time  `tpch:"o_orderdate"`
	OrderPriority string     `tpch:"o_orderpriority,15"`
	Clerk         string     `tpch:"o_clerk,15"`
	ShipPriority  int64      `tpch:"o_shippriority"`
	Comment       string     `tpch:"o_comment,79"`
}

//

var _ = _registerEntity(reflectEntity(
	rfl.Type[LineItem](),
	"line_item",
	[]string{"l_orderkey", "l_linenumber"},
	[][]string{{"l_partkey"}, {"l_suppkey"}},
))

type LineItem struct {
	OrderKey      Identifier `tpch:"l_orderkey"`
	PartKey       Identifier `tpch:"l_partkey"`
	SupplierKey   Identifier `tpch:"l_suppkey"`
	LineNumber    int64      `tpch:"l_linenumber"`
	Quantity      float64    `tpch:"l_quantity"`
	ExtendedPrice float64    `tpch:"l_extendedprice"`
	Discount      float64    `tpch:"l_discount"`
	Tax           float64    `tpch:"l_tax"`
	ReturnFlag    string     `tpch:"l_returnflag,1"`
	LineStatus    string     `tpch:"l_linestatus,1"`
	ShipDate      time.Time  `tpch:"l_shipdate"`
	CommitDate    time.Time  `tpch:"l_commitdate"`
	ReceiptDate   time.Time  `tpch:"l_receiptdate"`
	ShipInstruct  string     `tpch:"l_shipinstruct,25"`
	ShipMode      string     `tpch:"l_shipmode,10"`
	Comment       string     `tpch:"l_comment,44"`
}
