package main

import (
	"database/sql"
	"fmt"
	"go-todo/internal/server"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	host := "localhost"
	port := 5432
	user := "go_user"
	password := "go_password"
	dbname := "go_app_db"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	mux := http.NewServeMux()
	server.NewServer(mux, db)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
