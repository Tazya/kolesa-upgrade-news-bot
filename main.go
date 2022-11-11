package main

import (
	"net/http"

	"kolesa-upgrade/bot-messenger/http/handlers"
)

func main() {
	handler_enter := http.HandlerFunc(handlers.HandleRequest_enter)
	http.Handle("/", handler_enter)

	handler_health := http.HandlerFunc(handlers.HandleRequest_health)
	http.Handle("/health", handler_health)

	http.ListenAndServe(":8888", nil)
}
