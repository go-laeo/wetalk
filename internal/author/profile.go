package author

import (
	"context"

	"github.com/go-laeo/wetalk/ent"
)

type ProfileData struct {
	Name      string
	Account   string
	AvatarURL string
	Intro     string
	ID        int
}

func Profile(_ context.Context, u *ent.User) (*ProfileData, error) {
	data := &ProfileData{
		ID:        u.ID,
		Account:   u.Account,
		Name:      u.Name,
		AvatarURL: u.AvatarURL,
		Intro:     u.Intro,
	}

	return data, nil
}
