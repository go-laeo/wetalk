package auth

import (
	"context"

	"github.com/go-laeo/wetalk/ent"
	"github.com/go-laeo/wetalk/ent/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginData struct {
	Account  string
	Password string
}

// Login checks password is right.
func Login(ctx context.Context, data *LoginData) (*ent.User, error) {
	db := ent.FromContext(ctx)
	u, err := db.User.Query().Where(user.Account(data.Account)).First(ctx)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(data.Password))
	if err != nil {
		return nil, err
	}

	return u, nil
}
