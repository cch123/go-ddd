package protocol

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// InitHTTPServer init http server
func InitHTTPServer() {
	r := chi.NewRouter()

	r.Route("/order/v1", func(r chi.Router) {
		r.Get("/", listBooks)
		r.Get("/{id}", getBook)
	})
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

func listBooks(wr http.ResponseWriter, r *http.Request) {
	println("list books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	println("get book")
}
