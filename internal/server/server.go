package server

import (
	"github.com/Atlasp/bookcase/internal/core/domain/book"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	Router  chi.Router
	Service *book.Service
}

func New(service *book.Service) *Router {
	r := chi.NewRouter()
	s := &Router{
		Router:  r,
		Service: service,
	}

	s.addRoutes()

	return s
}
