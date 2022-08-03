package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
  r := chi.NewRouter()

  r.Use(middleware.Heartbeat("/ping"))
  r.Use(middleware.Logger)
  r.Use(middleware.CleanPath)

  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte("root"))
  })

  http.ListenAndServe(":3000", r)
}

