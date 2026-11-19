package main

import (
	"errors"
	"log"
	"net/http"
)

var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	ctx := r.Context()
	//Get the Session Token from the Cookie
	st, err := r.Cookie("session_token")

	if err != nil || st.Value == "" {
		log.Println("session_token is not set ", err)
		return AuthError
	}

	//check db for cookie
	var found int
	query := `SELECT 1 FROM users WHERE session_token = $1`
	err = db.QueryRow(ctx, query, st.Value).Scan(&found)
	if err != nil {
		log.Printf("User not found: %v\n", err)
		return AuthError
	}

	return nil
}
