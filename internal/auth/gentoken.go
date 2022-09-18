package auth

import (
	"context"
	"strconv"
	"time"

	"github.com/cristalhq/jwt/v4"
	"github.com/go-laeo/wetalk/ent"
	"github.com/go-laeo/wetalk/internal/config"
)

func GenToken(ctx context.Context, u *ent.User) (*jwt.Token, error) {
	cfg := config.FromContext(ctx)

	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(u.ID),
		Audience:  jwt.Audience{"wetalk"},
		ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "wetalk",
	}

	sig, err := jwt.NewSignerHS(jwt.HS384, []byte(cfg.SecretToken))
	if err != nil {
		return nil, err
	}
	tok, err := jwt.NewBuilder(sig).Build(claims)
	if err != nil {
		return nil, err
	}

	return tok, nil
}
