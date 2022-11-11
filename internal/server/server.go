package server

import (
	"kolesa-upgrade-team/delivery-bot/pkg/delievery"
	"net/http"
)

func NewServer(handler delievery.Handler) *http.Server {
	mux := http.NewServeMux()
	handler.InitRoutes(mux)
	return &http.Server{
		Addr:    ":8888",
		Handler: mux,
	}
}
