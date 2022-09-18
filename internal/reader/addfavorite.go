package reader

import (
	"context"

	"github.com/go-laeo/wetalk/ent"
)

func AddFavorite(ctx context.Context, u *ent.User, p *ent.Post) error {
	tx, err := ent.FromContext(ctx).Tx(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = tx.User.UpdateOne(u).AddFavoritePosts(p).Exec(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
