package web

import (
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter(handler *Handler) *Router {
	r := &Router{
		mux: http.NewServeMux(),
	}
	r.mux.Handle("/game/", handler)
	return r
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
