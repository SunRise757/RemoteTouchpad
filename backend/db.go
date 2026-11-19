package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn
var ctx context.Context = context.Background()

func initDb() {
	//ALTER ROLE "admin" SUPERUSER CREATEDB CREATEROLE INHERIT LOGIN REPLICATION BYPASSRLS;

	//this should be in .env file
	dbUrl := "postgres://admin:password@localhost:5432/remote_touchpad"

	var err error
	db, err = pgx.Connect(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	_, err = db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			session_token VARCHAR(255),
			csrf_token VARCHAR(255)
		)
	`)
	//add test user
	hashedPassword, _ := hashPassword("password")
	_, err = db.Exec(ctx, `
		INSERT INTO users (id, username, password)
		VALUES (1, 'admin', $1)
		ON CONFLICT (username) DO NOTHING
	`, hashedPassword)

	if err != nil {
		log.Fatalf("Unable to create users table: %v\n", err)
	}
}
