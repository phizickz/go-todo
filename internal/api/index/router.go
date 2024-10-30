package index

import (
	"net/http"
)

func (c *IndexController) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/", http.HandlerFunc(c.ServeIndex))
}
