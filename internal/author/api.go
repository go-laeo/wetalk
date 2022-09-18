package author

import (
	"github.com/go-laeo/pi"
	"github.com/go-laeo/wetalk/internal/auth"
)

func HTTPProfile() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		u := auth.FromContext(ctx.Context())
		data, err := Profile(ctx.Context(), u)
		if err != nil {
			return err
		}

		return ctx.Json(data)
	}
}

func HTTPPublish() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		p := &PublishData{}
		err := pi.Format(ctx, p)
		if err != nil {
			return err
		}

		u := auth.FromContext(ctx.Context())
		err = Publish(ctx.Context(), u, p)
		if err != nil {
			return pi.NewError(400, err.Error())
		}

		return pi.NewError(201, "created")
	}
}

func HTTPListGroup() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		gg, err := ListGroup(ctx.Context())
		if err != nil {
			return pi.NewError(500, err.Error())
		}

		return ctx.Json(gg)
	}
}

func HTTPUpdateProfile() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		p := &ProfileData{}
		err := pi.Format(ctx, p)
		if err != nil {
			return err
		}

		u := auth.FromContext(ctx.Context())
		err = UpdateProfile(ctx.Context(), u, p)
		if err != nil {
			return pi.NewError(500, err.Error())
		}

		return pi.NewError(204, "update successful!")
	}
}
