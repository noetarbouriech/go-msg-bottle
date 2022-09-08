package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

var tokenAuth *jwtauth.JWTAuth

var users = map[string]string{
	// username : password
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Input error"))
		return
	}

	// check username and password
	expectedPassword, ok := users[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Wrong password or username"))
		return
	}

	// create token
	expireTime := time.Now().Add(10 * time.Minute)
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
		"name": credentials.Username, // username
		"iat":  time.Now(),           // issued time
		"exp":  expireTime.Unix(),    // expire time
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error with token creation"))
		return
	}

	// put token in client cookies
	http.SetCookie(w, &http.Cookie{
		Name:       "jwt",
		Value:      tokenString,
		Path:       "",
		Domain:     os.Getenv("FRONTEND_DOMAIN"),
		Expires:    expireTime,
		RawExpires: "",
		MaxAge:     10000,
		Secure:     false,
		HttpOnly:   true,
		SameSite:   http.SameSiteLaxMode,
		Raw:        "",
		Unparsed:   []string{},
	})

	w.Write([]byte("Connection successful"))
	fmt.Println(time.Now().UTC().String() + " - " + credentials.Username + " connected")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials

	// decode json
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Input error"))
		return
	}

	// check if user already exists
	_, exists := users[credentials.Username]
	if exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User already exists"))
		return
	}

	// check if username is too short
	if len(credentials.Username) <= 3 || len(credentials.Username) >= 20 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username too short or too long"))
		return
	}

	// check if password is too short
	if len(credentials.Password) <= 6 || len(credentials.Password) >= 40 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Password not complex enough"))
		return
	}

	// add user to map
	users[credentials.Username] = credentials.Password

	w.Write([]byte("Account created"))
	fmt.Println(time.Now().UTC().String() + " - User " + credentials.Username + " has been created")
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	b := new(bytes.Buffer)
	for user, password := range users {
		fmt.Fprintf(b, "%s = \"%s\"\n", user, password)
	}
	w.Write(b.Bytes())
}
