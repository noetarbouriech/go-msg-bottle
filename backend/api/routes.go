package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func PublicRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root"))
	})

	r.Post("/login", Login)
	r.Post("/signup", SignUp)
}

func AdminRoutes(r chi.Router) {
	// authentification
	r.Use(jwtauth.Verifier(tokenAuth))
	r.Use(jwtauth.Authenticator)

	r.Get("/users", ListUsers)
	r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		w.Write([]byte(fmt.Sprintf("hello %v", claims["name"])))
	})
}
