package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{Router: mux.NewRouter()}
}

func (h *Handler) SetupRoutes() {
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Server is up")
	})
}