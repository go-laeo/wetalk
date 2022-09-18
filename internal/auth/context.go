package auth

import (
	"context"

	"github.com/go-laeo/wetalk/ent"
)

type _ctxkey struct{}

func NewContext(ctx context.Context, u *ent.User) context.Context {
	return context.WithValue(ctx, &_ctxkey{}, u)
}

func FromContext(ctx context.Context) *ent.User {
	if v := ctx.Value(&_ctxkey{}); v != nil {
		return v.(*ent.User)
	}
	return nil
}
