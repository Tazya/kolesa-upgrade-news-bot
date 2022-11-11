package delievery

import (
	"html/template"
	"net/http"
)

type Handler struct {
	tmpl *template.Template
}

func NewHandler() *Handler {
	return &Handler{
		tmpl: template.Must(template.ParseGlob("./ui/templates/*html")),
	}
}

func (h *Handler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/health", h.Health)
	mux.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("./ui/"))))
}
