package author

import (
	"context"
	"fmt"

	"github.com/go-laeo/wetalk/ent"
)

type PublishData struct {
	Content string
	GroupID int
}

func Publish(ctx context.Context, u *ent.User, data *PublishData) (err error) {
	tx, err := ent.FromContext(ctx).Tx(ctx)
	if err != nil {
		return fmt.Errorf("start TX failed: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	ctx = ent.NewContext(ctx, tx.Client()) // propagates the tx

	txc := tx.Post.Create().SetAuthor(u).SetContent(data.Content)
	if data.GroupID > 0 {
		g, err := tx.Group.Get(ctx, data.GroupID)
		if err != nil {
			return err
		}
		txc.SetGroup(g)
	}
	err = txc.Exec(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
