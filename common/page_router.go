package common

import (
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

type RenderFunc = func(name string, data PageData, status int)

type PageFunc = func(p RenderFunc, r *http.Request)
type PageData = map[string]any

type PageRouter struct {
	*mux.Router
}

func (m *PageRouter) Page(path string, p PageFunc) *mux.Route {
	return m.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		p(func(name string, data PageData, status int) {
			m.RenderPage(w, name, data)
		}, r)
	})
}

func NewPageRouter(f fs.FS) *PageRouter {
	// pagesFs := os.DirFS("pages")
	ParsePages(f)

	return &PageRouter{
		Router: mux.NewRouter(),
	}
}
