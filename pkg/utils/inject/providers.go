package inject

type Provider interface {
	provide() any
}
