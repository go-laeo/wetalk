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
	GroupSet  []Group
	ID        int
}

func Profile(ctx context.Context, u *ent.User) (*ProfileData, error) {
	data := &ProfileData{
		ID:        u.ID,
		Account:   u.Account,
		Name:      u.Name,
		AvatarURL: u.AvatarURL,
		Intro:     u.Intro,
	}

	rows, err := u.QueryGroups().All(ctx)
	if err != nil {
		return nil, err
	}

	for _, g := range rows {
		data.GroupSet = append(data.GroupSet, Group{
			ID:    g.ID,
			Name:  g.Name,
			Intro: g.Intro,
		})
	}

	return data, nil
}
