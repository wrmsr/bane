package c

//

type TypeSpecifier interface {
	DeclarationSpecifier
	isTypeSpecifier()
}

type typeSpecifier struct {
	node
}

func (t typeSpecifier) isTypeSpecifier() {}

//

type TypedefTypeSpecifier struct {
	typeSpecifier
	I Identifier
}

//

type BuiltinType interface {
	isBuiltinType()

	Name() string
}

type builtinType struct {
	typeSpecifier
}

func (t builtinType) isBuiltinType() {}

//

type Void struct {
	builtinType
}

type Char struct {
	builtinType
}

type Short struct {
	builtinType
}

type Int struct {
	builtinType
}

type Long struct {
	builtinType
}

type Float struct {
	builtinType
}

type Double struct {
	builtinType
}

type Signed struct {
	builtinType
}

type Unsigned struct {
	builtinType
}

type Bool struct {
	builtinType
}

type Complex struct {
	builtinType
}

type M128 struct {
	builtinType
}

type M128d struct {
	builtinType
}

type M128i struct {
	builtinType
}

func (t Void) Name() string     { return "void" }
func (t Char) Name() string     { return "char" }
func (t Short) Name() string    { return "short" }
func (t Int) Name() string      { return "int" }
func (t Long) Name() string     { return "long" }
func (t Float) Name() string    { return "float" }
func (t Double) Name() string   { return "double" }
func (t Signed) Name() string   { return "signed" }
func (t Unsigned) Name() string { return "unsigned" }
func (t Bool) Name() string     { return "_Bool" }
func (t Complex) Name() string  { return "_Complex" }
func (t M128) Name() string     { return "__m128" }
func (t M128d) Name() string    { return "__m128d" }
func (t M128i) Name() string    { return "__m128i" }

func ParseBuiltinType(s string) BuiltinType {
	switch s {
	case "void":
		return Void{}
	case "char":
		return Char{}
	case "short":
		return Short{}
	case "int":
		return Int{}
	case "long":
		return Long{}
	case "float":
		return Float{}
	case "double":
		return Double{}
	case "signed":
		return Signed{}
	case "unsigned":
		return Unsigned{}
	case "_Bool":
		return Bool{}
	case "_Complex":
		return Complex{}
	case "__m128":
		return M128{}
	case "__m128d":
		return M128d{}
	case "__m128i":
		return M128i{}
	}
	panic(s)
}

//

type BuiltinTypeSpecifier struct {
	typeSpecifier
	T BuiltinType
}

//

type StorageClassSpecifier int8

const (
	DefaultStorage StorageClassSpecifier = iota
	TypedefStorage
	ExternStorage
	StaticStorage
	ThreadLocalStorage
	AutoStorage
	RegisterStorage
)

func (s StorageClassSpecifier) String() string {
	switch s {
	case DefaultStorage:
		return "<default>"
	case TypedefStorage:
		return "typedef"
	case ExternStorage:
		return "extern"
	case StaticStorage:
		return "static"
	case ThreadLocalStorage:
		return "_Thread_local"
	case AutoStorage:
		return "auto"
	case RegisterStorage:
		return "register"
	}
	panic(s)
}

func ParseStorageClassSpecifier(s string) StorageClassSpecifier {
	switch s {
	case "typedef":
		return TypedefStorage
	case "extern":
		return ExternStorage
	case "static":
		return StaticStorage
	case "_Thread_local":
		return ThreadLocalStorage
	case "auto":
		return AutoStorage
	case "register":
		return RegisterStorage
	}
	panic(s)
}

//

type TypeQualifier int8

const (
	DefaultType TypeQualifier = iota
	ConstType
	RestrictType
	VolatileType
	AtomicType
)

func (q TypeQualifier) String() string {
	switch q {
	case DefaultType:
		return "<default>"
	case ConstType:
		return "const"
	case RestrictType:
		return "restrict"
	case VolatileType:
		return "volatile"
	case AtomicType:
		return "_Atomic"
	}
	panic(q)
}

func ParseTypeQualifier(s string) TypeQualifier {
	switch s {
	case "const":
		return ConstType
	case "restrict":
		return RestrictType
	case "volatile":
		return VolatileType
	case "_Atomic":
		return AtomicType
	}
	panic(s)
}

//

type SpecifierQualifier struct {
	Ts TypeSpecifier
	Tq TypeQualifier
}

type StructOrUnionSpecifier struct {
	typeSpecifier
	I  string
	Ds []StructDeclaration
}
