package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		array := []string{"1", "b"}
		pos := r.URL.Query().Get("index")
		i, err := strconv.Atoi(pos)
		if err != nil {
			panic(err)
		}
		log.Println("array[i]", array[i])
		str := "hello from get " + array[i]
		w.Write([]byte(str))
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		array := []string{"1", "b"}
		pos := r.URL.Query().Get("index")
		i, err := strconv.Atoi(pos)
		if err != nil {
			panic(err)
		}
		str := "hello from post " + array[i]
		w.Write([]byte(str))
	})

	log.Println("server started running on port 5555")
	http.ListenAndServe(":5555", r)
}
