package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Get("/api/vast", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("vast"))
	})
	r.Post("/api/rtb", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("rtb"))
	})
	r.Post("/api/js", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("js"))
	})
	r.Post("/api/hb", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hb"))
	})

	log.Println("server started running on port 5555")
	http.ListenAndServe(":5555", r)
}
