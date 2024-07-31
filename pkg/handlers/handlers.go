package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

var htmxRoot string = "web/html/"
var templateRoot string = "web/templates/"

type Card struct {
	Title string
	Body  string
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(templateRoot+"layout.html", templateRoot+"index.html"))
	tmpl.Execute(w, nil)
}

func ApiTesting(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	fmt.Println(r.Form)
}

func Api(w http.ResponseWriter, r *http.Request) {
	requestPath := strings.TrimPrefix(r.URL.Path, "/api/")
	if requestPath == "" {
		http.NotFound(w, r)
		return
	}

	methodHandlers := map[string]func(w http.ResponseWriter, r *http.Request, rPath string){
		"GET":  ApiGet,
		"POST": ApiPost,
		"PUT":  ApiPost,
	}

	if handler, found := methodHandlers[r.Method]; found {
		handler(w, r, requestPath)
	} else {
		http.Error(w, "Method not supported.", http.StatusMethodNotAllowed)
	}
}

func ApiPost(w http.ResponseWriter, r *http.Request, rPath string) {
	safePath := filepath.Clean(rPath)
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

func ApiGet(w http.ResponseWriter, r *http.Request, rPath string) {
	safePath := filepath.Clean(rPath)

	filePath := filepath.Join(htmxRoot, safePath+".html")
	http.ServeFile(w, r, filePath)
}
