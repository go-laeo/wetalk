package reader

import (
	"context"

	"github.com/go-laeo/wetalk/ent"
)

type CommentData struct {
	Content string
}

func AddComment(ctx context.Context, p *ent.Post, u *ent.User, data *CommentData) error {
	db := ent.FromContext(ctx)
	return db.Comment.Create().SetAuthor(u).SetContent(data.Content).SetPost(p).Exec(ctx)
}
