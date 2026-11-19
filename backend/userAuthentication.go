package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRTFToken     string
}

var registerTokens = map[string]int64{}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//get username and password from request
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid request method", err)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	//	registerToken := r.FormValue("registerToken")
	//	fmt.Println("token = ", registerToken)

	//check db for username and password
	query := `SELECT password FROM users WHERE username = $1`

	var DBpassword string = ""
	err := db.QueryRow(ctx, query, username).Scan(&DBpassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		log.Fatalf("Failed to execute query %v\n", err)
		return
	}
	if !checkPasswordHash(password, DBpassword) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	//generate and set session session_token and csrf_token
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
	query = `UPDATE users SET session_token=$1, csrf_token=$2 WHERE username=$3`
	_, err = db.Exec(ctx, query, sessionToken, csrfToken, username)
	if err != nil {
		log.Fatalf("Failed to execute query %v\n", err)
	}

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
	ctx = r.Context()
	if err := Authorize(r); err != nil {
		err := http.StatusUnauthorized
		http.Error(w, "Unauthorized", err)
		return
	}
	session_token := r.FormValue("session_token")

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
	query := `UPDATE users SET session_token = NULL, csrf_token = NULL WHERE session_token=$1`
	_, err := db.Exec(ctx, query, session_token)
	if err != nil {
		log.Fatalf("Failed to execute query %v\n", err)
	}

	fmt.Fprintln(w, "Logged out successfully!")

}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	ctx = r.Context()
	log.Println("register")
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", err)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	registerToken := r.FormValue("registerToken")

	if tokenTime, ok := registerTokens[registerToken]; !ok || tokenTime+5*60 < time.Now().Unix() {
		err := http.StatusUnauthorized
		http.Error(w, "invalid register token", err)
		return
	}
	if len(username) < 3 || len(password) < 8 {
		err := http.StatusNotAcceptable
		http.Error(w, "invalid username/password", err)
		return
	}

	//check if username already exists
	query := `SELECT username FROM users WHERE username = $1`
	rows, err := db.Query(ctx, query, username)
	if err != nil {
		log.Println("1")
		log.Fatalf("Failed to execute query %v\n", err)
	}
	defer rows.Close()
	if rows.Next() {
		err := http.StatusConflict
		http.Error(w, "User already exists", err)
		return
	}

	//add user and password to db
	query = `
	INSERT INTO users (username, password)
	VALUES ($1, $2)
	`
	hashedPassword, _ := hashPassword(password)
	_, err = db.Exec(ctx, query, username, hashedPassword)
	if err != nil {
		log.Println("2")
		log.Fatalf("Failed to execute query %v\n", err)
	}

	delete(registerTokens, registerToken)
	fmt.Println("User registred successfully!")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":  true,
		"message":  "Register successful",
		"rediredt": "login",
	})
}
