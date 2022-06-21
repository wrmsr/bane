package dot

import "fmt"

//

func typeError(o any) error {
	return fmt.Errorf("%T", o)
}

//

type Item interface {
	isItem()
}

type item struct{}

func (i item) isItem() {}

//

type Value interface {
	Item
	isValue()
}

type value struct {
	item
}

func (v value) isValue() {}

func AsValue(o any) Value {
	if o, ok := o.(Value); ok {
		return o
	}
	switch o := o.(type) {
	case string:
		return Text{Text: o}
	case []any:
		return AsTable(o)
	}
	panic(typeError(o))
}

//

type Raw struct {
	value
	Raw string
}

func AsRaw(o any) Raw {
	switch o := o.(type) {
	case Raw:
		return o
	case string:
		return Raw{Raw: o}
	}
	panic(typeError(o))
}

//

type Text struct {
	value
	Text string
}

func AsText(o any) Text {
	switch o := o.(type) {
	case Text:
		return o
	case string:
		return Text{Text: o}
	}
	panic(typeError(o))
}

//

type Cell struct {
	item
	Value Value
}

func AsCell(o any) Cell {
	switch o := o.(type) {
	case Cell:
		return o
	}
	return Cell{Value: AsValue(o)}
}

//

type Row struct {
	item
	Cells []Cell
}

func AsRow(o any) Row {
	switch o := o.(type) {
	case Row:
		return o
	case []any:
		cs := make([]Cell, len(o))
		for i, c := range o {
			cs[i] = AsCell(c)
		}
		return Row{Cells: cs}
	}
	panic(typeError(o))
}

//

type Table struct {
	value
	Rows []Row
}

func AsTable(o any) Table {
	switch o := o.(type) {
	case Table:
		return o
	case []any:
		rs := make([]Row, len(o))
		for i, r := range o {
			rs[i] = AsRow(r)
		}
		return Table{Rows: rs}
	}
	panic(typeError(o))
}

//

type Id struct {
	item
	Id string
}

func AsId(o any) Id {
	switch o := o.(type) {
	case Id:
		return o
	case string:
		return Id{Id: o}
	}
	panic(typeError(o))
}

//

type Attrs struct {
	item
	Attrs map[string]Value
}

func AsAttrs(o any) Attrs {
	switch o := o.(type) {
	case Attrs:
		return o
	case map[string]any:
		m := make(map[string]Value, len(o))
		for k, v := range o {
			m[k] = AsValue(v)
		}
		return Attrs{Attrs: m}
	}
	panic(typeError(o))
}

//

type Stmt interface {
	isStmt()
}

type stmt struct{}

func (s stmt) isStmt() {}

//

type RawStmt struct {
	stmt
	Raw string
}

func AsRawStmt(o any) RawStmt {
	switch o := o.(type) {
	case RawStmt:
		return o
	case string:
		return RawStmt{Raw: o}
	}
	panic(typeError(o))
}

//

type Edge struct {
	stmt

	Left  Id
	Right Id

	Attrs Attrs
}

//

type Node struct {
	stmt

	Id Id

	Attrs Attrs
}

//

type Graph struct {
	Stmts []Stmt

	Id Id
}
