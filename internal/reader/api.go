package reader

import (
	"strconv"

	"github.com/go-laeo/pi"
	"github.com/go-laeo/wetalk/ent"
	"github.com/go-laeo/wetalk/internal/auth"
)

func HTTPCreateFavorite() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		u := auth.FromContext(ctx.Context())

		db := ent.FromContext(ctx.Context())

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return pi.NewError(404, "post not found")
		}

		p, err := db.Post.Get(ctx.Context(), id)
		if err != nil {
			return pi.NewError(404, "post not found")
		}

		err = AddFavorite(ctx.Context(), u, p)
		if err != nil {
			return pi.NewError(500, err.Error())
		}

		return pi.NewError(201, "created")
	}
}

func HTTPCreateComment() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		p := &CommentData{}
		err := pi.Format(ctx, p)
		if err != nil {
			return err
		}

		u := auth.FromContext(ctx.Context())

		db := ent.FromContext(ctx.Context())

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return pi.NewError(404, "post not found")
		}

		po, err := db.Post.Get(ctx.Context(), id)
		if err != nil {
			return pi.NewError(404, "post not found")
		}

		err = AddComment(ctx.Context(), po, u, p)
		if err != nil {
			return pi.NewError(500, err.Error())
		}

		return pi.NewError(201, "created")
	}
}
