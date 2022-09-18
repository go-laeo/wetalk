package auth

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/cristalhq/jwt/v4"
	"github.com/go-laeo/pi"
	"github.com/go-laeo/wetalk/ent"
	"github.com/go-laeo/wetalk/internal/config"
)

func FromToken(ctx context.Context, token string) (*ent.User, error) {
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		return nil, pi.NewError(401, "token is empty")
	}

	cfg := config.FromContext(ctx)
	sig, err := jwt.NewSignerHS(jwt.HS384, []byte(cfg.SecretToken))
	if err != nil {
		return nil, pi.NewError(500, err.Error())
	}

	rc := &jwt.RegisteredClaims{}
	err = jwt.ParseClaims([]byte(token), sig, rc)
	if err != nil {
		return nil, pi.NewError(401, err.Error())
	}

	if !rc.IsValidAt(time.Now()) {
		return nil, pi.NewError(401, "session expired")
	}

	id, err := strconv.Atoi(rc.Subject)
	if err != nil {
		return nil, pi.NewError(401, "unauthorized")
	}

	db := ent.FromContext(ctx)
	u, err := db.User.Get(ctx, id)
	if err != nil {
		return nil, pi.NewError(401, "user not found")
	}

	return u, nil
}
