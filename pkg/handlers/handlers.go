package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var htmxRoot string = "web/html/"
var templateRoot string = "web/templates/"

func ApiGet(w http.ResponseWriter, r *http.Request, rPath string) {
	safePath := filepath.Clean(rPath)

	filePath := filepath.Join(htmxRoot, safePath+".html")
	http.ServeFile(w, r, filePath)
}

func ApiTesting(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	fmt.Println(r.Form)
}

