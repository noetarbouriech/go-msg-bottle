package main

import (
	"fmt"
	"net/http"

	"github.com/noetarbouriech/go-msg-bottle/api"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
  r.Use(cors.Handler(cors.Options{
  	AllowedOrigins: []string{"https://*", "http://*"},
  	AllowOriginFunc: func(r *http.Request, origin string) bool {
  		return true
  	},
  	AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
  	AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
  	ExposedHeaders:     []string{"Set-Cookie"},
  	AllowCredentials:   true,
  	MaxAge:             300,
  	OptionsPassthrough: false,
  	Debug:              false,
  }))

	// routes
	r.Group(api.PublicRoutes)
	r.Group(api.AdminRoutes)

	http.ListenAndServe(":3000", r)
}
