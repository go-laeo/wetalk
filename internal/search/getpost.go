package search

import (
	"context"

	"github.com/go-laeo/wetalk/ent"
	"github.com/go-laeo/wetalk/ent/post"
)

func GetPost(ctx context.Context, id int, u *ent.User) (*Post, error) {
	db := ent.FromContext(ctx)
	po, err := db.Post.Query().Where(post.ID(id)).WithAuthor().WithGroup().First(ctx)
	if err != nil {
		return nil, err
	}

	p := Post{}
	err = FromEntPost(ctx, &p, po, u)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
