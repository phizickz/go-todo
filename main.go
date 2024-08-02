package main

import (
	"log"
	"net/http"

	"go-todo/pkg/handlers"
	"go-todo/pkg/models/card"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	http.Handle("/test", http.NotFoundHandler())
	http.HandleFunc("/test/", handlers.ApiTesting)
	
	http.Handle("/api", http.NotFoundHandler())
	http.HandleFunc("/api/card", handlers.CardApi)
	card.NewCard("abc","def")
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
