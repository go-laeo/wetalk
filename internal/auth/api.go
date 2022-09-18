package auth

import (
	"github.com/go-laeo/pi"
	"github.com/go-laeo/wetalk/ent"
)

type TokenData struct {
	Token string
}

func HTTPLogin() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		p := &LoginData{}
		err := pi.Format(ctx, p)
		if err != nil {
			return err
		}

		u, err := Login(ctx.Context(), p)
		if err != nil && !ent.IsNotFound(err) {
			return pi.NewError(400, "Password incorrect")
		}
		if ent.IsNotFound(err) {
			return pi.NewError(400, "Account not found")
		}

		// Generate token for u
		at, err := GenToken(ctx.Context(), u)
		if err != nil {
			return pi.NewError(400, err.Error())
		}

		return ctx.Json(TokenData{Token: at.String()})
	}
}

func HTTPCreateAccount() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		p := &AccountData{}
		err := pi.Format(ctx, p)
		if err != nil {
			return err
		}

		err = CreateAccount(ctx.Context(), p)
		if err != nil {
			return pi.NewError(400, err.Error())
		}

		return ctx.Json(pi.NewError(201, "created"))
	}
}

func HTTPUpdatePassword() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		p := &PasswordData{}
		err := pi.Format(ctx, p)
		if err != nil {
			return err
		}

		u := FromContext(ctx.Context())
		err = UpdatePassword(ctx.Context(), u, p)
		if err != nil {
			return pi.NewError(400, err.Error())
		}

		return ctx.Json(pi.NewError(204, "updated"))
	}
}
