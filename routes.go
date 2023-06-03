package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/yusuf/mailapp/handlers"
)

func Routes(lg handlers.Logic) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RealIP)

	mux.Get("/", lg.Home())
	mux.Post("/api/submit", lg.GetSubscriber())
	mux.Post("/api/send", lg.SendMail())

	return mux
}
