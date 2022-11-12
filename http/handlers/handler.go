package handlers

import (
	"net/http"
)

func InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/health", Health)
	mux.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("./ui/"))))
}
