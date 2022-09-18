package config

import "context"

type _ctxkey struct{}

func NewContext(ctx context.Context, c *Config) context.Context {
	return context.WithValue(ctx, &_ctxkey{}, c)
}

func FromContext(ctx context.Context) *Config {
	if v := ctx.Value(&_ctxkey{}); v != nil {
		return v.(*Config)
	}
	return nil
}
