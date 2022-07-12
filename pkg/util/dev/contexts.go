package dev

import "context"

func WithDefaultTimeout(ctx context.Context) (context.Context, func()) {
	return context.WithTimeout(ctx, ProvideAs[Config]().DefaultTimeout)
}

func DefaultTimeout() (context.Context, func()) {
	return WithDefaultTimeout(context.Background())
}
