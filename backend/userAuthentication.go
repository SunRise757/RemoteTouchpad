package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRTFToken     string
}

var users = map[string]Login{}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request = ")
	fmt.Println(r)
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid request method", err)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, ok := users[username]
	if !ok || !checkPasswordHash(password, user.HashedPassword) {
		err := http.StatusUnauthorized
		http.Error(w, "Invalid username or password", err)
		return
	}

	sessionToken := generateToken(32)
	csrfToken := generateToken(32)

	//Set session Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	//Set CSRF token in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false, //Needs to be accessible to the client-side
	})

	//Store tokens in the database
	user.SessionToken = sessionToken
	user.CSRTFToken = csrfToken
	users[username] = user

	fmt.Println("Login successful")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"message":  "Login successful",
		"rediredt": "touchpad",
	})
}

func protectedRouteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)

	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "invalid request method", err)
		return
	}

	if err := Authorize(r); err != nil {
		err := http.StatusUnauthorized
		http.Error(w, "Unauthorized", err)
		return
	}

	username := r.FormValue("username")
	fmt.Fprintf(w, "CSRF validation successfil! Welcome, %s", username)

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if err := Authorize(r); err != nil {
		err := http.StatusUnauthorized
		http.Error(w, "Unauthorized", err)
		return
	}

	// Clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	//Clear the tokens from the database
	username := r.FormValue("username")
	user, _ := users[username]
	user.SessionToken = ""
	user.CSRTFToken = ""
	users[username] = user

	fmt.Fprintln(w, "Logged out successfully!")

}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", err)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if len(username) < 3 || len(password) < 8 {
		err := http.StatusNotAcceptable
		http.Error(w, "invalid username/password", err)
		return
	}
	if _, ok := users[username]; ok {
		err := http.StatusConflict
		http.Error(w, "User already exists", err)
		return
	}
	hashedPassword, _ := hashPassword(password)
	users[username] = Login{
		HashedPassword: hashedPassword,
	}

	fmt.Fprintln(w, "User registred successfully!")
}
