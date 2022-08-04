package main

import (
	"fmt"
	"net/http"

	"github.com/noetarbouriech/go-msg-bottle/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	fmt.Println("Starting server")
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
  r.Use(middleware.AllowContentType("application/json"))
  r.Use(render.SetContentType(render.ContentTypeJSON))

	// routes
	r.Group(api.PublicRoutes)
	r.Group(api.AdminRoutes)

	http.ListenAndServe(":3000", r)
}
