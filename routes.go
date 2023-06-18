package main

import (
	"net/http"

	"github.com/akinbyte/mailapp/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(lg handlers.Logic) *chi.Mux {
	mux := chi.NewRouter()

	// default middleware to logger requests and avoid server shutting down
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	// setting up endpoints to respective handlers
	mux.Get("/", lg.Home())
	mux.Post("/api/submit", lg.GetSubscriber())
	mux.Post("/api/send", lg.SendMail())

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", lg.Home())
	// mux.HandleFunc("/api/submit", lg.GetSubscriber())
	// mux.HandleFunc("/api/send", lg.SendMail())

	return mux
}
