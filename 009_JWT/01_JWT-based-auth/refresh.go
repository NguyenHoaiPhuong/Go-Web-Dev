package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Refresh : handler function
func Refresh(w http.ResponseWriter, r *http.Request) {
	log.Println("Refresh")

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

	// (END) The code up-till this point is the same as the first part of the `Welcome` route

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	if claims.ExpiresAt-time.Now().Unix() > 30 {
		log.Println(claims.ExpiresAt)
		log.Println(time.Now().Unix())
		err := errors.New("A new token will only be issued if the old token is within 30 seconds of expiry")
		ResponseError(err, http.StatusBadRequest, w)
		return
	}

	// Now, create a new token for the current use, with a renewed expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	if err != nil {
		ResponseError(err, http.StatusInternalServerError, w)
		return
	}

	// Set the new token as the users `token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
