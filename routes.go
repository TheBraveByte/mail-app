package main

import (
	_"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/yusuf/mailapp/handlers"
)

func Routes(lg handlers.Logic) *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Get("/", lg.Home())
	mux.Post("/api/submit", lg.GetSubscriber())
	mux.Post("/api/send", lg.SendMail())

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", lg.Home())
	// mux.HandleFunc("/api/submit", lg.GetSubscriber())
	// mux.HandleFunc("/api/send", lg.SendMail())

	return mux
}
