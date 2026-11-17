package main

import (
	"errors"
	"log"
	"net/http"
)

var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	//TEST ONLY this thing needs to be in db
	username := "admin"
	//username := r.FormValue("username")
	user, ok := users[username]

	if !ok {
		log.Println("username not found")
		return AuthError
	}

	//Get the Session Token from the Cookie
	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		log.Println("session_token check failed ", err, " ", st.Value, " ", user.SessionToken)
		return AuthError
	}

	//Get the CSRF token from the headers
	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRTFToken || csrf == "" {
		log.Println("csrf_token check failed", csrf, " ", user.CSRTFToken)
		return AuthError
	}

	return nil
}
