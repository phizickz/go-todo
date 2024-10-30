package index

import (
	"go-todo/web/views"
	"net/http"

	"github.com/a-h/templ"
)

type IndexController struct{}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (c *IndexController) ServeIndex(w http.ResponseWriter, r *http.Request) {
	handler := templ.Handler(views.Index())
	handler.ServeHTTP(w, r)
}
