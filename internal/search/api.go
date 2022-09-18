package search

import (
	"strconv"

	"github.com/go-laeo/pi"
	"github.com/go-laeo/wetalk/internal/auth"
)

func HTTPListPost() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		p, _ := strconv.Atoi(ctx.Query("pi", "1"))
		ps, _ := strconv.Atoi(ctx.Query("ps", "20"))
		if p < 1 {
			p = 1
		}
		if ps < 10 || ps > 20 {
			ps = 20
		}

		u, _ := auth.FromToken(ctx.Context(), ctx.Get("authorization"))

		rows, err := ListTopPost(ctx.Context(), u, (p-1)*ps, ps)
		if err != nil {
			return pi.NewError(500, err.Error())
		}
		return ctx.Json(rows)
	}
}

func HTTPQueryPost() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return pi.NewError(404, "post not found")
		}

		u, _ := auth.FromToken(ctx.Context(), ctx.Get("authorization"))

		p, err := GetPost(ctx.Context(), id, u)
		if err != nil {
			return pi.NewError(404, err.Error())
		}

		return ctx.Json(p)
	}
}

func HTTPQueryCommentList() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return pi.NewError(404, "post not found")
		}

		p, _ := strconv.Atoi(ctx.Query("pi", "1"))
		ps, _ := strconv.Atoi(ctx.Query("ps", "20"))
		if p < 1 {
			p = 1
		}
		if ps < 10 || ps > 20 {
			ps = 20
		}

		rows, err := GetCommentList(ctx.Context(), id, (p-1)*ps, ps)
		if err != nil {
			return pi.NewError(500, err.Error())
		}

		return ctx.Json(rows)
	}
}
