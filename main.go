package main

import (
	"net/http"

	"github.com/Atlasp/bookcase/internal/core/adapters"
	"github.com/Atlasp/bookcase/internal/core/domain/book"
	"github.com/Atlasp/bookcase/internal/server"
)

func main() {
	repo := adapters.NewMemMap()
	service := book.NewService(repo)

	r := server.New(service)

	http.ListenAndServe("localhost:8080", r.Router)
}
