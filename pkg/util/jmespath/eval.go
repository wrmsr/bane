package jmespath

/*
class Evaluator(ta.Generic[T], dispatch.Class):

    def __init__(self, runtime: Runtime[T]) -> None:
        super().__init__()

        self._runtime = runtime

    __call__ = dispatch.property()

    def __call__(self, node: n.Node, obj: T) -> T:  # type: ignore  # noqa
        raise TypeError(node)

    def __call__(self, node: n.And, obj: T) -> T:  # type: ignore  # noqa
        left = self(node.left, obj)
        if self._runtime.is_truthy(left):
            return self(node.right, obj)
        else:
            return left

    def __call__(self, node: n.Compare, obj: T) -> T:  # type: ignore  # noqa
        left = self(node.left, obj)
        right = self(node.right, obj)
        return self._runtime.compare(node.op, left, right)

    def __call__(self, node: n.CreateArray, obj: T) -> T:  # type: ignore  # noqa
        if self._runtime.is_null(obj):
            return obj
        else:
            return self._runtime.create_array([self(child, obj) for child in node.items])

    def __call__(self, node: n.CreateObject, obj: T) -> T:  # type: ignore  # noqa
        if self._runtime.is_null(obj):
            return obj
        else:
            return self._runtime.create_object({field: self(child, obj) for field, child in node.fields.items()})

    def __call__(self, node: n.Current, obj: T) -> T:  # type: ignore  # noqa
        return obj

    def __call__(self, node: n.ExpressionRef, obj: T) -> T:  # type: ignore  # noqa
        return self(node.expr, obj)

    def __call__(self, node: n.FlattenArray, obj: T) -> T:  # type: ignore  # noqa
        if self._runtime.get_type(obj) == ValueType.ARRAY:
            lst = []
            for item in self._runtime.to_iterable(obj):
                if self._runtime.get_type(item) == ValueType.ARRAY:
                    lst.extend(self._runtime.to_iterable(item))
                else:
                    lst.append(item)
            return lst
        else:
            return self._runtime.create_null()

    def __call__(self, node: n.FlattenObject, obj: T) -> T:  # type: ignore  # noqa
        if self._runtime.get_type(obj) == ValueType.OBJECT:
            return self._runtime.create_array(self._runtime.to_iterable(obj))
        else:
            return self._runtime.create_null()

    def __call__(self, node: n.FunctionCall, obj: T) -> T:  # type: ignore  # noqa
        args = []
        for arg in node.args:
            if isinstance(arg, n.ExpressionRef):
                args.append(NodeArg(arg))
            else:
                args.append(ValueArg(self(arg, obj)))
        return self._runtime.invoke_function(node.name, args)

    def __call__(self, node: n.Index, obj: T) -> T:  # noqa
        if self._runtime.get_type(obj) == ValueType.ARRAY:
            items = self._runtime.to_iterable(obj)
            i = node.value
            if i < 0:
                i = len(items) + 1
            if 0 <= i < len(items):
                return items[i]
        return self._runtime.create_null()

    def __call__(self, node: n.JsonLiteral, obj: T) -> T:  # type: ignore  # noqa
        return self._runtime.parse_str(node.text)

    def __call__(self, node: n.Negate, obj: T) -> T:  # type: ignore  # noqa
        return self._runtime.create_bool(self._runtime.is_truthy(self(node.item, obj)))

    def __call__(self, node: n.Or, obj: T) -> T:  # type: ignore  # noqa
        left = self(node.left, obj)
        if self._runtime.is_truthy(left):
            return left
        else:
            return self(node.right, obj)

    def __call__(self, node: n.Project, obj: T) -> T:  # type: ignore  # noqa
        if self._runtime.get_type(obj) == ValueType.ARRAY:
            items = [
                oitem
                for iitem in self._runtime.to_iterable(obj)
                for oitem in [self(node.child, iitem)]
                if not self._runtime.is_null(oitem)
            ]
            return self._runtime.create_array(items)
        else:
            return self._runtime.create_null()

    def __call__(self, node: n.Property, obj: T) -> T:  # type: ignore  # noqa
        return self._runtime.get_property(obj, node.name)

    def __call__(self, node: n.Selection, obj: T) -> T:  # type: ignore  # noqa
        if self._runtime.get_type(obj) == ValueType.ARRAY:
            items = [
                item
                for item in self._runtime.to_iterable(obj)
                if self._runtime.is_truthy(node.child, item)
            ]
            return self._runtime.create_array(items)
        else:
            return self._runtime.create_null()

    def __call__(self, node: n.Sequence, obj: T) -> T:  # type: ignore  # noqa
        for child in node.items:
            obj = self(child, obj)
        return obj

    def __call__(self, node: n.Slice, obj: T) -> T:  # type: ignore  # noqa
        items = self._runtime.to_iterable(obj)
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
        return self._runtime.create_array(lst)

    def __call__(self, node: n.String, obj: T) -> T:  # type: ignore  # noqa
        return self._runtime.create_str(node.value)

    def __call__(self, node: n.Parameter, obj: T) -> T:  # type: ignore  # noqa
        if isinstance(node.target, n.Parameter.NameTarget):
            return self._runtime.get_name_var(node.target.value)
        elif isinstance(node.target, n.Parameter.NumberTarget):
            return self._runtime.get_num_var(node.target.value)
        else:
            raise TypeError(node.target)
*/
