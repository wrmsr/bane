package def

//

type X_structExpectsDef struct {
	Name string

	Fields   []string
	Defaults []string

	Inits int
}

func (sed X_structExpectsDef) isRootDef() {}

func (sed X_structExpectsDef) Register() any {
	addPackageRootDef(globalRegistry, getCallingPackage(), sed)
	return nil
}

//

type X_structInit struct {
	Defaults map[string]any

	Inits []any
}

//

type structInitBuilder struct {
	reg *packageRegistry
}

func (b *structInitBuilder) buildInit() map[string]X_structInit {
	return nil
}
