package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

var users = map[string]string{
  "user1": "password1",
  "user2": "password2",
}

type Credentials struct {
  Username string `json:"username"`
  Password string `json:"password"`
}

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
  var credentials Credentials

  // decode json
  err := json.NewDecoder(r.Body).Decode(&credentials)
  if err != nil {
		w.Write([]byte("Input error"))
    return
  }

  // check username/password
  expectedPassword, ok := users[credentials.Username]
  if !ok || expectedPassword != credentials.Password {
		w.Write([]byte("Wrong password or username"))
    return
  }

  // create token
	expireTime := time.Now().Add(10 * time.Minute)
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
		"name": credentials.Username, // username
    "iat": time.Now(),            // issued time
		"exp": expireTime.Unix(),     // expire time
	})
	if err != nil {
		w.Write([]byte("Error with token creation"))
		return
	}

  // put token in client cookies
	http.SetCookie(w,&http.Cookie{
		Name:       "jwt",
		Value:      tokenString,
		Path:       "",
		Domain:     "",
		Expires:    expireTime,
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   []string{},
	})

  w.Write([]byte("Connection successful"))
  fmt.Println(time.Now().UTC().String() + " - " + credentials.Username + " connected")
}

// TODO
func SignUp(w http.ResponseWriter, r *http.Request) {
  // Add credentials to db
  // Login
}
