package consteval

import msh "github.com/wrmsr/bane/pkg/util/marshal"

var _ = msh.RegisterTo[Value](
	msh.SetImplOf[Array](),
	msh.SetImplOf[Basic](),
	msh.SetImplOf[Call](),
	msh.SetImplOf[Dynamic](),
	msh.SetImplOf[Map](),
	msh.SetImplOf[Nil](),
	msh.SetImplOf[Struct](),
	msh.SetImplOf[Type](),
)

var _ = msh.RegisterTo[BasicKind](
	msh.SetStringerEnumTypes(
		CharBasic,
		FloatBasic,
		ImagBasic,
		IntBasic,
		StringBasic,
	),
)

var _ = msh.RegisterTo[Dynamic](
	msh.SetField{Name: "a", Omit: true},
)
