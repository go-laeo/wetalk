package author

import (
	"context"

	"github.com/go-laeo/wetalk/ent"
)

type Group struct {
	Name  string
	Intro string
	ID    int
}

func ListGroup(ctx context.Context) ([]Group, error) {
	db := ent.FromContext(ctx)
	rows, err := db.Group.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	b := make([]Group, 0, len(rows))
	for _, g := range rows {
		b = append(b, Group{
			ID:    g.ID,
			Name:  g.Name,
			Intro: g.Intro,
		})
	}

	return b, nil
}
