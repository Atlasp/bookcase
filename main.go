package main

import (
	"log"
	"net/http"

	"github.com/Atlasp/bookcase/internal/adapters"
	"github.com/Atlasp/bookcase/internal/core/book"
	"github.com/Atlasp/bookcase/internal/server"
)

func main() {
	repo := adapters.NewMemMap()
	service := book.NewService(repo)

	r := server.New(service)

	log.Fatal(http.ListenAndServe("localhost:8080", r.Router))
}
