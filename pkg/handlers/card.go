package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

type Card struct {
	Title string
	Body  string
}

func CardApi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		PostCard(w, r)
	} else {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
}

func PostCard(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	requestPath := strings.TrimPrefix(r.URL.Path, "/api/")
	if requestPath == "" {
		http.NotFound(w, r)
		return
	}

	safePath := filepath.Clean(requestPath)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Create a new card instance
	card := Card{
		Title: r.FormValue("title"),
		Body:  r.FormValue("body"),
	}

	tmpl := template.Must(template.ParseFiles(filepath.Join(templateRoot, safePath+".html")))
	tmpl.Execute(w, card)
}
