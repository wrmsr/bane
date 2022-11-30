package c

import (
	msh "github.com/wrmsr/bane/pkg/util/marshal"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

var _ = msh.Register(rfl.TypeOf[Declaration](),
	msh.SetImplOf[ParameterDeclaration](),
	msh.SetImplOf[FunctionDeclaration](),
)
