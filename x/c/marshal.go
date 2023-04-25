package c

import (
	msh "github.com/wrmsr/bane/core/marshal"
)

var _ = msh.RegisterTo[Node](
	msh.InheritImplsOf[Declaration](),
	msh.InheritImplsOf[Expression](),
	msh.InheritImplsOf[Statement](),
	msh.InheritImplsOf[TypeSpecifier](),
	msh.SetImplOf[TranslationUnit](),
)

var _ = msh.RegisterTo[Declaration](
	msh.SetImplOf[DeclaratorsDeclaration](),
	msh.SetImplOf[FunctionDeclaration](),
	msh.SetImplOf[ParameterDeclaration](),
)

var _ = msh.RegisterTo[DeclarationSpecifier](
	msh.InheritImplsOf[TypeSpecifier](),
	msh.SetImplOf[TypeQualifier](),
)

var _ = msh.RegisterTo[TypeSpecifier](
	msh.SetImplOf[Char](),
	msh.SetImplOf[Int](),
	msh.SetImplOf[Long](),
	msh.SetImplOf[StructOrUnionSpecifier](),
)

var _ = msh.RegisterTo[Declarator](
	msh.SetImplOf[IdentifierDeclarator](),
	msh.SetImplOf[InitDeclarator](),
	msh.SetImplOf[ParameterListDeclarator](),
	msh.SetImplOf[PointerDeclarator](),
)

var _ = msh.RegisterTo[Statement](
	msh.SetImplOf[ExpressionStatement](),
	msh.SetImplOf[ReturnStatement](),
)

var _ = msh.RegisterTo[Expression](
	msh.SetImplOf[Binary](),
	msh.SetImplOf[Call](),
	msh.SetImplOf[Constant](),
	msh.SetImplOf[Identifier](),
	msh.SetImplOf[StringLiteral](),
)

var _ = msh.RegisterTo[BinaryOp](
	msh.SetStringerEnumTypes(
		AddOp,
		SubOp,
		MulOp,
		DivOp,
		ModOp,

		LtOp,
		LteOp,
		EqOp,
		NeOp,
		GtOp,
		GteOp,

		BitAndOp,
		BitOrOp,
		BitXorOp,
	),
)

var _ = msh.RegisterTo[TypeQualifier](
	msh.SetStringerEnumTypes(
		DefaultType,
		ConstType,
		RestrictType,
		VolatileType,
		AtomicType,
	),
)

var _ = msh.RegisterTo[BlockItem](
	msh.InheritImplsOf[Declaration](),
	msh.InheritImplsOf[Statement](),
)
