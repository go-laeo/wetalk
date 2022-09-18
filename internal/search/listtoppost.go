package search

import (
	"context"
	"time"

	"github.com/go-laeo/wetalk/ent"
	"github.com/go-laeo/wetalk/internal/author"
)

type Post struct {
	CreatedAt    time.Time
	Group        *author.Group
	Content      string
	Author       author.ProfileData
	ID           int
	Positived    bool
	Negatived    bool
	Favorited    bool
	CommentCount int
}

func ListTopPost(ctx context.Context, u *ent.User, offset, limit int) ([]Post, error) {
	db := ent.FromContext(ctx)
	rows, err := db.Post.Query().WithAuthor().WithGroup().Offset(offset).Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}

	b := make([]Post, 0, len(rows))
	for _, row := range rows {
		p := Post{}
		err = FromEntPost(ctx, &p, row, u)
		if err != nil {
			return nil, err
		}

		b = append(b, p)
	}

	return b, nil
}
