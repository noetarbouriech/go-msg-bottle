package main

import (
	"fmt"
	"net/http"

	"github.com/noetarbouriech/go-msg-bottle/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Starting server")
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)

	// routes
	r.Group(api.PublicRoutes)
	r.Group(api.AdminRoutes)

	http.ListenAndServe(":3000", r)
}
