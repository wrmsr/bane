package inject

type BindingsRegistry struct {
}

func (r *BindingsRegistry) Bind(o ...any) *BindingsRegistry {
	return r
}
