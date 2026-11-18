package main

import (
	"errors"
	"log"
	"net/http"
)

var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	//Get the Session Token from the Cookie
	st, err := r.Cookie("session_token")

	if err != nil || st.Value == "" {
		log.Println("session_token is not set ", err)
		return AuthError
	}

	for username, user := range users {
		if user.SessionToken == st.Value {
			log.Println(username, "authorised")
			return nil
		}
	}

	log.Println("session_token check failed ", err)
	return AuthError

	/* websockets cant have custom headers
	//Get the CSRF token from the headers
	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRTFToken || csrf == "" {
		log.Println("csrf_token check failed", csrf, " ", user.CSRTFToken)
		return AuthError
	}

	*/

}
