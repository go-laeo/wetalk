package auth

import (
	"context"
	"fmt"

	"github.com/go-laeo/wetalk/ent"
	"golang.org/x/crypto/bcrypt"
)

type PasswordData struct {
	Password  string
	Successor string
}

func UpdatePassword(ctx context.Context, u *ent.User, data *PasswordData) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(data.Password))
	if err != nil {
		return fmt.Errorf("password incorrect")
	}

	b, err := bcrypt.GenerateFromPassword([]byte(data.Successor), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("generate secret failed: %w", err)
	}

	db := ent.FromContext(ctx)
	return db.User.UpdateOne(u).SetPassword(string(b)).Exec(ctx)
}
