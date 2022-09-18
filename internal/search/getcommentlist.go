package search

import (
	"context"
	"time"

	"github.com/go-laeo/wetalk/ent"
	"github.com/go-laeo/wetalk/ent/comment"
	"github.com/go-laeo/wetalk/ent/post"
	"github.com/go-laeo/wetalk/internal/author"
)

type Comment struct {
	CreatedAt time.Time
	Content   string
	Author    author.ProfileData
	ID        int
}

func GetCommentList(ctx context.Context, id int, offset, limit int) ([]Comment, error) {
	db := ent.FromContext(ctx)
	rows, err := db.Comment.Query().
		Where(comment.HasPostWith(post.ID(id))).
		Order(ent.Desc(comment.FieldID)).
		WithAuthor().
		Offset(offset).Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}

	b := make([]Comment, 0, len(rows))
	for _, row := range rows {
		c := Comment{}
		err = FromEntComment(ctx, row, &c, nil)
		if err != nil {
			return nil, err
		}

		b = append(b, c)
	}

	return b, nil
}
