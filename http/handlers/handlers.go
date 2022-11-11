package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func HandleRequest_health(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status": "ok",
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func HandleRequest_enter(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := os.ReadFile("Hello.jpg")
	if err != nil {
		panic(err)
	}
	w.Write(fileBytes)
}
