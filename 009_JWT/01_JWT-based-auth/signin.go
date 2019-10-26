package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Signin : handle function for signin
func Signin(w http.ResponseWriter, r *http.Request) {
	log.Println("Signin")

	var creds Credentials

	// Decode the content of r.Body and save into creds
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		ResponseError(err, http.StatusBadRequest, w)
		return
	}

	if !IsCredentialsValid(creds) {
		err = errors.New("username or password missmatched")
		ResponseError(err, http.StatusUnauthorized, w)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	// At this stage, token has header and payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string by the given secret key. Now, token got enough info: header, payload and signature
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		ResponseError(err, http.StatusInternalServerError, w)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
