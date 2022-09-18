package author

import (
	"context"

	"github.com/go-laeo/wetalk/ent"
)

func UpdateProfile(ctx context.Context, u *ent.User, data *ProfileData) error {
	db := ent.FromContext(ctx)
	return db.User.UpdateOne(u).SetName(data.Name).SetIntro(data.Intro).Exec(ctx)
}
