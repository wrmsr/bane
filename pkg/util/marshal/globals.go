package marshal

import "reflect"

//

var globalRegistry = NewRegistry(nil)

func GlobalRegistry() *Registry {
	return globalRegistry
}

func Register(ty reflect.Type, items ...RegistryItem) *Registry {
	return globalRegistry.Register(ty, items...)
}

//

var globalManager = NewDefaultManager(func() *Registry { return globalRegistry })

func GlobalManager() *Manager {
	return globalManager
}

func Marshal(v any, o ...MarshalOpt) (Value, error) {
	return globalManager.Marshal(v, o...)

}

func Unmarshal(mv Value, v any, o ...UnmarshalOpt) error {
	return globalManager.Unmarshal(mv, v, o...)
}
