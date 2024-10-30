package index

import (
	"go-todo/web/views"
	"net/http"

	"github.com/a-h/templ"
)

func (c *IndexController) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/", templ.Handler(views.Index()))
}
