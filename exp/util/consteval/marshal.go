package consteval

import msh "github.com/wrmsr/bane/pkg/util/marshal"

var _ = msh.RegisterTo[Value](
	msh.SetImplOf[Basic](),
	msh.SetImplOf[Type](),
	msh.SetImplOf[Struct](),
	msh.SetImplOf[Array](),
	msh.SetImplOf[Map](),
	msh.SetImplOf[Dynamic](),
)

var _ = msh.RegisterTo[BasicKind](
	msh.SetStringerEnumTypes(
		IntBasic,
		FloatBasic,
		ImagBasic,
		CharBasic,
		StringBasic,
	),
)

var _ = msh.RegisterTo[Dynamic](
	msh.SetField{Name: "a", Omit: true},
)
