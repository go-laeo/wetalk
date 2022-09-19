package force

import (
	"context"
	"errors"

	"github.com/go-laeo/wetalk/ent"
	"github.com/go-laeo/wetalk/ent/coin"
	"github.com/go-laeo/wetalk/ent/group"
	"github.com/go-laeo/wetalk/ent/user"
)

const (
	PriceCreateGroup = 10240
)

var (
	ErrInsufficientBalance   = errors.New("insufficient balance")
	ErrGroupNameAlreadyTaken = errors.New("group name already taken")
)

type GroupData struct {
	Name  string
	Intro string
}

func CreateGroup(ctx context.Context, u *ent.User, data *GroupData) error {
	tx, err := ent.FromContext(ctx).Tx(ctx)
	if err != nil {
		return err
	}
	defer tx.Commit()

	dup, err := tx.Group.Query().Where(group.Name(data.Name)).Exist(ctx)
	if err != nil {
		return err
	}
	if dup {
		return ErrGroupNameAlreadyTaken
	}

	c, err := tx.Coin.Query().Where(coin.HasOwnerWith(user.ID(u.ID))).Order(ent.Desc(coin.FieldID)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrInsufficientBalance
		}
		return err
	}
	if c.Balance-PriceCreateGroup < 0 {
		return ErrInsufficientBalance
	}

	g, err := tx.Group.Create().SetName(data.Name).SetIntro(data.Intro).Save(ctx)
	if err != nil {
		return err
	}

	c, err = tx.Coin.Create().SetOwner(u).
		SetDeal("CreateGroup " + g.Name).
		SetAmount(-PriceCreateGroup).SetBalance(c.Balance - PriceCreateGroup).
		Save(ctx)
	if err != nil {
		return err
	}

	err = tx.User.UpdateOne(u).SetCoin(c.Balance).Exec(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
