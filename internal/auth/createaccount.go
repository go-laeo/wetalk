package auth

import (
	"context"
	"errors"

	"github.com/go-laeo/wetalk/ent"
	"github.com/go-laeo/wetalk/ent/user"
	"golang.org/x/crypto/bcrypt"
)

type AccountData struct {
	Name     string
	Account  string
	Password string
}

func CreateAccount(ctx context.Context, data *AccountData) (err error) {
	tx, err := ent.FromContext(ctx).Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	dup, err := tx.User.Query().Where(user.Account(data.Account)).Exist(ctx)
	if err != nil {
		return err
	}
	if dup {
		return errors.New("account already exists")
	}

	buf, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u, err := tx.User.Create().
		SetAccount(data.Account).SetName(data.Name).
		SetPassword(string(buf)).
		Save(ctx)
	if err != nil {
		return err
	}

	cnt, err := tx.User.Query().Count(ctx)
	if err != nil {
		return err
	}
	if cnt == 1 {
		err = tx.Coin.Create().SetDeal("创世空投").SetAmount(1000000).SetBalance(1000000).SetOwner(u).Exec(ctx)
		if err != nil {
			return err
		}
		err = u.Update().SetCoin(1000000).Exec(ctx)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
