package server

import (
	"fmt"
	"net/http"

	"github.com/Atlasp/bookcase/internal/core/domain/book"
	"github.com/pingcap/log"
)

func (rt *Router) AddBook() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b := book.Book{
			Title:   "To Kill A Mocking Bird",
			Author:  "Harper Lee",
			Edition: 12,
			Pages:   309,
			Read:    false,
		}
		err := rt.Service.AddBook(b)
		if err != nil {
			log.Error("cannot add book")
		}
		fmt.Fprintf(w, "%s added to the bookshelf", b.Title)
	}
}
