package json

import (
	msh "github.com/wrmsr/bane/pkg/util/marshal"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

var _ = msh.Register(rfl.TypeOf[Node](),
	msh.SetImplOf[Object]("obj"),
	msh.SetImplOf[Array]("array"),
	msh.SetImplOf[String]("string"),
	msh.SetImplOf[Number]("number"),
	msh.SetImplOf[True]("true"),
	msh.SetImplOf[False]("false"),
	msh.SetImplOf[Null]("null"),
)
