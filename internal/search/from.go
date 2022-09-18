package search

import (
	"context"

	"github.com/go-laeo/wetalk/ent"
	"github.com/go-laeo/wetalk/ent/user"
	"github.com/go-laeo/wetalk/internal/author"
)

func FromEntPost(ctx context.Context, p *Post, po *ent.Post, u *ent.User) (err error) {
	p.ID = po.ID
	p.Content = po.Content
	p.CreatedAt = po.CreatedAt

	if po.Edges.Author != nil {
		p.Author = author.ProfileData{
			ID:        po.Edges.Author.ID,
			Name:      po.Edges.Author.Name,
			Account:   po.Edges.Author.Account,
			AvatarURL: po.Edges.Author.AvatarURL,
			Intro:     po.Edges.Author.Intro,
		}
	}

	if po.Edges.Group != nil {
		p.Group = &author.Group{
			ID:   po.Edges.Group.ID,
			Name: po.Edges.Group.Name,
		}
	}

	if u != nil {
		cnt, err := po.QueryFavoriteUsers().Where(user.ID(u.ID)).Count(ctx)
		if err != nil {
			return err
		}
		if cnt > 0 {
			p.Favorited = true
		}
	}

	p.CommentCount, err = po.QueryComments().Count(ctx)
	return err
}

func FromEntComment(_ context.Context, co *ent.Comment, c *Comment, _ *ent.User) error {
	c.ID = co.ID
	c.Content = co.Content
	c.CreatedAt = co.CreatedAt

	if co.Edges.Author != nil {
		c.Author = author.ProfileData{
			ID:        co.Edges.Author.ID,
			Name:      co.Edges.Author.Name,
			Account:   co.Edges.Author.Account,
			AvatarURL: co.Edges.Author.AvatarURL,
		}
	}

	return nil
}
