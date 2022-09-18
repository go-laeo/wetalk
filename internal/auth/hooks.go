package auth

import (
	"github.com/go-laeo/pi"
)

func HookHTTPAuth(next pi.HandlerFunc) pi.HandlerFunc {
	return func(ctx pi.Context) error {
		u, err := FromToken(ctx.Context(), ctx.Get("authorization"))
		if err != nil {
			return err
		}

		ctx.SetContext(NewContext(ctx.Context(), u))

		return next(ctx)
	}
}
