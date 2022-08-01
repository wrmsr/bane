package json

import (
	msh "github.com/wrmsr/bane/pkg/util/marshal"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

var _ = msh.Register(rfl.TypeOf[Node](),
	msh.SetImplOf[Object](),
	msh.SetImplOf[Array](),
	msh.SetImplOf[String](),
	msh.SetImplOf[Number](),
	msh.SetImplOf[True](),
	msh.SetImplOf[False](),
	msh.SetImplOf[Null](),
)
