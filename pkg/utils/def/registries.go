package def

import (
	"errors"

	rtu "github.com/wrmsr/bane/pkg/utils/runtime"
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
		panic(errors.New("must not call from this package"))
	}
	return pkg
}

//

type packageRegistry struct {
	name string
}

func newPackageRegistry(name string) *packageRegistry {
	return &packageRegistry{
		name: name,
	}
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

func (r *registry) Package(name string) *packageRegistry {
	if p, ok := r.packages[name]; ok {
		return p
	}
	p := newPackageRegistry(name)
	r.packages[name] = p
	return p
}
