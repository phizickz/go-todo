package main

import (
	"go-todo/internal/server"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	app := server.NewServer()
	app.RegisterRoutes(mux)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
