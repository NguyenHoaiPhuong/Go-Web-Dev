package main

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Credentials : Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Claims : Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func IsCredentialsValid(creds Credentials) bool {
	username := creds.Username
	password := creds.Password
	if pass, ok := users[username]; !ok || password != pass {
		return false
	}
	return true
}

// ResponseError : response error to http.ResponseWriter
func ResponseError(err error, status int, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}

// ResponseMessage : write message to http.ResponseWriter
func ResponseMessage(msg string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}
