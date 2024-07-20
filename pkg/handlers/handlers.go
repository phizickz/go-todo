package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/index.html"))
	tmpl.Execute(w, nil)
}

func Api(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ApiGet(w, r)
	}

	http.Error(w, "Method is not supported.", http.StatusNotFound)
}

func ApiGet(w http.ResponseWriter, r *http.Request) {
	requestPath := strings.TrimPrefix(r.URL.Path, "/api/html/")
	if requestPath == "" {
		http.NotFound(w, r)
		return
	}

	safePath := filepath.Clean(requestPath)
	htmxRoot := "web/html/"

	filePath := filepath.Join(htmxRoot, safePath+".html")
	http.ServeFile(w, r, filePath)
}
