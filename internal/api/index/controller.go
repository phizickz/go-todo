package index

import "net/http"

type IndexController struct{}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (c *IndexController) ServeIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<div>HELLO</div>"))
}
