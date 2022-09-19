package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/go-laeo/pi"
	"github.com/go-laeo/wetalk/ent"
	"github.com/go-laeo/wetalk/internal/auth"
	"github.com/go-laeo/wetalk/internal/author"
	"github.com/go-laeo/wetalk/internal/config"
	"github.com/go-laeo/wetalk/internal/force"
	"github.com/go-laeo/wetalk/internal/meta"
	"github.com/go-laeo/wetalk/internal/reader"
	"github.com/go-laeo/wetalk/internal/search"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	cc := config.FromEnv()
	ctx = config.NewContext(ctx, cc)

	ec, err := ent.Open("mysql", cc.DatabaseURL)
	if err != nil {
		panic(err)
	}

	err = ec.Schema.Create(ctx, schema.WithForeignKeys(false))
	if err != nil {
		panic(err)
	}

	ctx = ent.NewContext(ctx, ec.Debug())

	sm := pi.NewServerMux()
	sm.Route("/api/v1/tokens").Post(auth.HTTPLogin())
	sm.Route("/api/v1/users").Post(auth.HTTPCreateAccount())

	sm.Group("/api/v1", func(sm pi.ServerMux) {
		sm.Use(auth.HookHTTPAuth)

		sm.Route("/profile/password").Post(auth.HTTPUpdatePassword())
		sm.Route("/profile").
			Get(author.HTTPProfile()).
			Put(author.HTTPUpdateProfile())
		sm.Route("/posts").Post(author.HTTPPublish())
		sm.Route("/posts/:id/comments").Post(reader.HTTPCreateComment())
		sm.Route("/posts/:id/favorite_users").Post(reader.HTTPCreateFavorite())
		sm.Route("/groups").Post(force.HTTPCreateGroup())
	})

	sm.Route("/api/v1/groups").Get(author.HTTPListGroup())
	sm.Route("/api/v1/posts").Get(search.HTTPListPost())
	sm.Route("/api/v1/posts/:id").Get(search.HTTPQueryPost())
	sm.Route("/api/v1/posts/:id/comments").Get(search.HTTPQueryCommentList())
	sm.Route("/api/v1/meta").Get(meta.HTTPQueryMeta())

	sm.Route("/*static").Any(pi.FileServer(pi.Sub(http.FS(stubs), "web/dist"), "index.html"))

	srv := http.Server{
		Addr:    "localhost:8080",
		Handler: sm,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	log.Println("start listen on localhost:8080")
	go srv.ListenAndServe()

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	<-ctx.Done()
	cancel()

	log.Println("graceful shutdown will start after 10s")
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	srv.Shutdown(ctx)
	cancel()

	log.Println("bye")
}
