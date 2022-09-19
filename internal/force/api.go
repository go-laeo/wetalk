package force

import (
	"github.com/go-laeo/pi"
	"github.com/go-laeo/wetalk/internal/auth"
)

func HTTPCreateGroup() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		p := &GroupData{}
		err := pi.Format(ctx, p)
		if err != nil {
			return err
		}

		u := auth.FromContext(ctx.Context())

		err = CreateGroup(ctx.Context(), u, p)
		if err != nil {
			return pi.NewError(400, err.Error())
		}

		return pi.NewError(201, "created")
	}
}
