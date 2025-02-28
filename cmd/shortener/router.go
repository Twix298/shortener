package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/Twix298/shortener/internal/app/handlers"
)

func newRouter(h *handlers.BaseHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/", h.GetShortUrl)
	r.Get("/{id}", h.GetFullUrl)
	return r
}
