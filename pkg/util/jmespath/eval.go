package jmespath

import "fmt"

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
	}
	panic(fmt.Errorf("unhandled node type: %#v", node))
}

/*
   def __call__(self, node: n.Selection, obj: T) -> T:  # type: ignore  # noqa
       if e.rt.get_type(obj) == ValueType.ARRAY:
           items = [
               item
               for item in e.rt.to_iterable(obj)
               if e.rt.is_truthy(node.child, item)
           ]
           return e.rt.create_array(items)
       else:
           return e.rt.create_null()

   def __call__(self, node: n.Sequence, obj: T) -> T:  # type: ignore  # noqa
       for child in node.items:
           obj = self(child, obj)
       return obj

   def __call__(self, node: n.Slice, obj: T) -> T:  # type: ignore  # noqa
       items = e.rt.to_iterable(obj)
       step = node.step or 1
       rounding = (step + 1) if (step < 0) else (step - 1)
       limit = -1 if (step < 0) else 0
       start = node.start or limit
       stop = node.stop or (-2**32 if step < 0 else 2**32)
       begin = max(len(items) + start, 0) if (start < 0) else min(start, len(items) + limit)
       end = max(len(items) + stop, limit) if (stop < 0) else min(stop, len(items))
       steps = max(0, (end - begin + rounding) // step)
       lst = []
       i = 0
       offset = begin
       while i < steps:
           lst.append(items[offset])
           offset += step
       return e.rt.create_array(lst)

   def __call__(self, node: n.String, obj: T) -> T:  # type: ignore  # noqa
       return e.rt.create_str(node.value)

   def __call__(self, node: n.Parameter, obj: T) -> T:  # type: ignore  # noqa
       if isinstance(node.target, n.Parameter.NameTarget):
           return e.rt.get_name_var(node.target.value)
       elif isinstance(node.target, n.Parameter.NumberTarget):
           return e.rt.get_num_var(node.target.value)
       else:
           raise TypeError(node.target)
*/
