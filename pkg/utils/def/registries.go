package def

import (
	"errors"
	"reflect"

	rtu "github.com/wrmsr/bane/pkg/utils/runtime"
)

//

type RegistryError struct {
	err error
}

func (e RegistryError) Error() string {
	return e.err.Error()
}

func (e RegistryError) Unwrap() error {
	return e.err
}

var (
	RegistrySealedError = RegistryError{errors.New("registry sealed")}
)

//

func _getCallingPackage() string {
	sf := rtu.GetStackTrace(1, 3)[0]
	return rtu.ParseName(sf.Name).Pkg
}

var _thisPackage = func() string { return _getCallingPackage() }()

func getCallingPackage() string {
	pkg := _getCallingPackage()
	if pkg == _thisPackage {
		panic(RegistryError{errors.New("must not call from this package")})
	}
	return pkg
}

//

type X_packageInit struct {
	Structs map[string]X_structInit
}

//

type packageRegistry struct {
	name   string
	sealed bool

	rootDefs    []RootDef
	rootDefsMap map[reflect.Type][]RootDef

	init *X_packageInit
}

func newPackageRegistry(name string) *packageRegistry {
	return &packageRegistry{
		name: name,

		rootDefsMap: make(map[reflect.Type][]RootDef),
	}
}

func (r *packageRegistry) addRootDef(def RootDef) RootDef {
	if r.sealed {
		panic(RegistrySealedError)
	}

	r.rootDefs = append(r.rootDefs, def)

	dty := reflect.TypeOf(def)
	r.rootDefsMap[dty] = append(r.rootDefsMap[dty], def)

	return def
}

func (r *packageRegistry) addRootDefs(defs ...RootDef) *packageRegistry {
	if r.sealed {
		panic(RegistrySealedError)
	}
	for _, def := range defs {
		r.addRootDef(def)
	}
	return r
}

func (r *packageRegistry) seal() *packageRegistry {
	if r.sealed {
		return r
	}

	r.init = &X_packageInit{
		Structs: (&structInitBuilder{reg: r}).buildInit(),
	}

	r.sealed = true
	return r
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

func addPackageRootDef[T RootDef](r *registry, pkg string, def T) T {
	r.getPackage(pkg).addRootDef(def)
	return def
}

func X_getPackageInit() *X_packageInit {
	pr := globalRegistry.getPackage(getCallingPackage())
	pr.seal()
	return pr.init
}
