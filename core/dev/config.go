//go:build !nodev

package dev

import "time"

type Config struct {
	DefaultTimeout time.Duration
}

func DefaultConfig() Config {
	return Config{
		DefaultTimeout: time.Duration(10) * time.Second,
	}
}
