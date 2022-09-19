package meta

import (
	"github.com/go-laeo/pi"
	"github.com/go-laeo/wetalk/internal/config"
)

type Meta struct {
	SiteName string
}

func HTTPQueryMeta() pi.HandlerFunc {
	return func(ctx pi.Context) error {
		cc := config.FromContext(ctx.Context())

		return ctx.Json(Meta{
			SiteName: cc.SiteName,
		})
	}
}
