package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

// Welcome : handle function
func Welcome(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome")

	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			ResponseError(err, http.StatusUnauthorized, w)
			return
		}
		// For any other type of error, return a bad request status
		ResponseError(err, http.StatusBadRequest, w)
		return
	}

	// Get the JWT string from the cookie
	tokenString := c.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ResponseError(err, http.StatusUnauthorized, w)
			return
		}
		ResponseError(err, http.StatusBadRequest, w)
		return
	}
	if !token.Valid {
		ResponseError(err, http.StatusUnauthorized, w)
		return
	}

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	ResponseMessage(fmt.Sprintf("Welcome %s!", claims.Username), w)
}
