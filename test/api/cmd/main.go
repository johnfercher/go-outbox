package main

import (
	"github.com/johnfercher/go-outbox/test/api/internal/config"
	"github.com/johnfercher/go-outbox/test/api/internal/mysql"
	"github.com/johnfercher/go-outbox/test/api/internal/repository"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg, err := config.Load(os.Args)
	if err != nil {
		panic(err)
	}

	db, err := mysql.Start(cfg.Mysql.Url, cfg.Mysql.Db, cfg.Mysql.User, cfg.Mysql.Password)
	if err != nil {
		panic(err)
	}

	repo := repository.New(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := repo.AddNew()
		if err != nil {
			w.Write([]byte("fail"))
			return
		}
		w.Write([]byte("success"))
		return
	})
	http.ListenAndServe(":3000", r)
}
