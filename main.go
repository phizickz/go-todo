package main

import (
	"log"
	"net/http"

	"go-todo/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	http.Handle("/api", http.NotFoundHandler())
	http.HandleFunc("/api/", handlers.Api)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
