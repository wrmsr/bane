package c

//

type TypeSpecifier interface {
	isTypeSpecifier()
}

type typeSpecifier struct {
	node
}

func (t typeSpecifier) isTypeSpecifier() {}

//

type SimpleType interface {
	isSimpleType()

	Name() string
}

type simpleType struct {
	typeSpecifier
}

func (t simpleType) isSimpleType() {}

//

type Void struct {
	simpleType
}

type Char struct {
	simpleType
}

type Short struct {
	simpleType
}

type Int struct {
	simpleType
}

type Long struct {
	simpleType
}

type Float struct {
	simpleType
}

type Double struct {
	simpleType
}

type Signed struct {
	simpleType
}

type Unsigned struct {
	simpleType
}

type Bool struct {
	simpleType
}

type Complex struct {
	simpleType
}

type M128 struct {
	simpleType
}

type M128d struct {
	simpleType
}

type M128i struct {
	simpleType
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

func ParseSimpleType(s string) SimpleType {
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
