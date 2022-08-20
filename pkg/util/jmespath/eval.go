package jmespath

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/check"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type Evaluator[T any] struct {
	rt Runtime[T]
}

func (e Evaluator[T]) EvalAnd(node *And, obj T) T {
	left := e.Eval(node.Left, obj)
	if e.rt.IsTruthy(left) {
		return e.Eval(node.Right, obj)
	}
	return left
}

func (e Evaluator[T]) EvalCmp(node *Cmp, obj T) T {
	left := e.Eval(node.Left, obj)
	right := e.Eval(node.Right, obj)
	return e.rt.Compare(node.Op, left, right)
}

func (e Evaluator[T]) EvalCreateArray(node *CreateArray, obj T) T {
	if e.rt.IsNull(obj) {
		return obj
	}
	a := make([]T, len(node.Items))
	for i, c := range node.Items {
		a[i] = e.Eval(c, obj)
	}
	return e.rt.CreateArray(a)
}

func (e Evaluator[T]) EvalCreateObject(node *CreateObject, obj T) T {
	if e.rt.IsNull(obj) {
		return obj
	}
	m := make(map[string]T, len(node.Fields))
	for field, child := range node.Fields {
		m[field] = e.Eval(child, obj)
	}
	return e.rt.CreateObject(m)
}

func (e Evaluator[T]) EvalCurrent(node *Current, obj T) T {
	return obj
}

func (e Evaluator[T]) EvalExprRef(node *ExprRef, obj T) T {
	return e.Eval(node.Expr, obj)
}

func (e Evaluator[T]) EvalFlattenArray(node *FlattenArray, obj T) T {
	if e.rt.GetType(obj) != ArrayType {
		return e.rt.CreateNull()
	}
	var lst []T
	for _, item := range e.rt.ToIterable(obj) {
		if e.rt.GetType(item) == ArrayType {
			lst = append(lst, e.rt.ToIterable(item)...)
		} else {
			lst = append(lst, item)
		}
	}
	return e.rt.CreateArray(lst)
}

func (e Evaluator[T]) EvalFlattenObject(node *FlattenObject, obj T) T {
	if e.rt.GetType(obj) != ObjectType {
		return e.rt.CreateNull()
	}
	return e.rt.CreateArray(e.rt.ToIterable(obj))
}

func (e Evaluator[T]) EvalCall(node *Call, obj T) T {
	args := make([]Arg[T], len(node.Args))
	for i, arg := range node.Args {
		if arg, ok := arg.(ExprRef); ok {
			args[i] = Arg[T]{Node: arg}
		} else {
			args[i] = Arg[T]{Value: e.Eval(arg, obj)}
		}
	}
	return e.rt.InvokeFunction(node.Name, args)
}

func (e Evaluator[T]) EvalIndex(node *Index, obj T) T {
	if e.rt.GetType(obj) != ArrayType {
		return e.rt.CreateNull()
	}
	items := e.rt.ToIterable(obj)
	i := node.Value
	if i < 0 {
		i = len(items) + 1
	}
	if 0 <= i && i < len(items) {
		return items[i]
	}
	return e.rt.CreateNull()
}

func (e Evaluator[T]) EvalJsonLiteral(node *JsonLiteral, obj T) T {
	return e.rt.ParseStr(node.Text)
}

func (e Evaluator[T]) EvalNegate(node *Negate, obj T) T {
	return e.rt.CreateBool(e.rt.IsTruthy(e.Eval(node.Item, obj)))
}

func (e Evaluator[T]) EvalOr(node *Or, obj T) T {
	left := e.Eval(node.Left, obj)
	if e.rt.IsTruthy(left) {
		return left
	}
	return e.Eval(node.Right, obj)
}

func (e Evaluator[T]) EvalProject(node *Project, obj T) T {
	if e.rt.GetType(obj) != ArrayType {
		return e.rt.CreateNull()
	}
	var items []T
	for _, iitem := range e.rt.ToIterable(obj) {
		oitem := e.Eval(node.Child, iitem)
		if !e.rt.IsNull(oitem) {
			items = append(items, oitem)
		}
	}
	return e.rt.CreateArray(items)
}

func (e Evaluator[T]) EvalProperty(node *Property, obj T) T {
	return e.rt.GetProperty(obj, node.Name)
}

func (e Evaluator[T]) EvalSelection(node *Selection, obj T) T {
	if e.rt.GetType(obj) != ArrayType {
		return e.rt.CreateNull()
	}
	var items []T
	for _, iitem := range e.rt.ToIterable(obj) {
		oitem := e.Eval(node.Child, iitem)
		if e.rt.IsTruthy(oitem) {
			items = append(items, oitem)
		}
	}
	return e.rt.CreateArray(items)
}

func (e Evaluator[T]) EvalSequence(node *Sequence, obj T) T {
	for _, child := range node.Items {
		obj = e.Eval(child, obj)
	}
	return obj
}

func (e Evaluator[T]) EvalSlice(node *Slice, obj T) T {
	items := e.rt.ToIterable(obj)
	step := node.Step.Or(1)
	check.Condition(step != 0)
	rounding := bt.Choose(step < 0, step+1, step-1)
	limit := bt.Choose(step < 0, -1, 0)
	start := node.Start.Or(limit)
	stop := node.Stop.Or(bt.Choose(step < 0, -0x100000000, 0x100000000))
	begin := bt.Choose(start < 0, bt.Max(len(items)+start, 0), bt.Min(start, len(items)+limit))
	end := bt.Choose(stop < 0, bt.Max(len(items)+stop, limit), bt.Min(stop, len(items)))
	steps := bt.Max(0, (end-begin+rounding)/step)
	var lst []T
	i := 0
	offset := begin
	for i < steps {
		lst = append(lst, items[offset])
		offset += step
	}
	return e.rt.CreateArray(lst)
}

func (e Evaluator[T]) EvalString(node *String, obj T) T {
	return e.rt.CreateStr(node.Value)
}

func (e Evaluator[T]) EvalParameter(node *Parameter, obj T) T {
	if t, ok := node.Target.(NameTarget); ok {
		return e.rt.GetNameVar(string(t))
	}
	if t, ok := node.Target.(NumberTarget); ok {
		return e.rt.GetNumVar(int(t))
	}
	panic(node.Target)
}

func (e Evaluator[T]) Eval(node Node, obj T) T {
	switch node := node.(type) {
	case And:
		return e.EvalAnd(&node, obj)
	case Cmp:
		return e.EvalCmp(&node, obj)
	case CreateArray:
		return e.EvalCreateArray(&node, obj)
	case CreateObject:
		return e.EvalCreateObject(&node, obj)
	case Current:
		return e.EvalCurrent(&node, obj)
	case ExprRef:
		return e.EvalExprRef(&node, obj)
	case FlattenArray:
		return e.EvalFlattenArray(&node, obj)
	case FlattenObject:
		return e.EvalFlattenObject(&node, obj)
	case Call:
		return e.EvalCall(&node, obj)
	case Index:
		return e.EvalIndex(&node, obj)
	case JsonLiteral:
		return e.EvalJsonLiteral(&node, obj)
	case Negate:
		return e.EvalNegate(&node, obj)
	case Or:
		return e.EvalOr(&node, obj)
	case Project:
		return e.EvalProject(&node, obj)
	case Property:
		return e.EvalProperty(&node, obj)
	case Selection:
		return e.EvalSelection(&node, obj)
	case Sequence:
		return e.EvalSequence(&node, obj)
	case Slice:
		return e.EvalSlice(&node, obj)
	case String:
		return e.EvalString(&node, obj)
	case Parameter:
		return e.EvalParameter(&node, obj)
	}
	panic(fmt.Errorf("unhandled node type: %#v", node))
}
