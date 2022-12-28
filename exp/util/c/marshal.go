package c

import (
	msh "github.com/wrmsr/bane/pkg/util/marshal"
)

var _ = msh.RegisterTo[Declaration](
	msh.SetImplOf[ParameterDeclaration](),
	msh.SetImplOf[FunctionDeclaration](),
)

var _ = msh.RegisterTo[DeclarationSpecifier](
	msh.InheritImplsOf[TypeSpecifier](),
)

var _ = msh.RegisterTo[TypeSpecifier](
	msh.SetImplOf[Int](),
)

var _ = msh.RegisterTo[Declarator](
	msh.SetImplOf[ParameterListDeclarator](),
)

var _ = msh.RegisterTo[BlockItem](
	msh.InheritImplsOf[Statement](),
)

var _ = msh.RegisterTo[Statement](
	msh.SetImplOf[ReturnStatement](),
)

var _ = msh.RegisterTo[Expression](
	msh.SetImplOf[Identifier](),
)
