package def

import (
	"errors"

	rtu "github.com/wrmsr/bane/core/runtime"
)

//

type RegistryError struct {
	err error
}

func (e RegistryError) Error() string { return e.err.Error() }
func (e RegistryError) Unwrap() error { return e.err }

var (
	RegistrySealedError = RegistryError{errors.New("registry sealed")}
)

//

var _thisPackage = func() string { return rtu.GetCaller(0).Pkg }()

func X_defPackageName() string {
	return _thisPackage
}

func getCallingPackage() string {
	pkg := rtu.GetCaller(2).Pkg
	if pkg == _thisPackage {
		panic(RegistryError{errors.New("must not call from this package")})
	}
	return pkg
}

//

type packageRegistry struct {
	name  string
	defs  []PackageDef
	_spec *PackageSpec
}

func newPackageRegistry(name string) *packageRegistry {
	return &packageRegistry{
		name: name,
	}
}

func (r *packageRegistry) addDef(def PackageDef) {
	if r._spec != nil {
		panic(RegistrySealedError)
	}
	r.defs = append(r.defs, def)
}

func (r *packageRegistry) spec() *PackageSpec {
	if r._spec == nil {
		r._spec = NewPackageSpec(r.name, r.defs)
	}
	return r._spec
}

//

type registry struct {
	packages map[string]*packageRegistry
}

func newRegistry() *registry {
	return &registry{
		packages: make(map[string]*packageRegistry),
	}
}

func (r *registry) getPackage(name string) *packageRegistry {
	if p, ok := r.packages[name]; ok {
		return p
	}
	p := newPackageRegistry(name)
	r.packages[name] = p
	return p
}

//

var globalRegistry = newRegistry()

func addPackageDef[T PackageDef](r *registry, pkg string, def T) T {
	r.getPackage(pkg).addDef(def)
	return def
}

func X_addPackageDef(def PackageDef) {
	globalRegistry.getPackage(getCallingPackage()).addDef(def)
}

func X_getPackageSpec() *PackageSpec {
	return globalRegistry.getPackage(getCallingPackage()).spec()
}
