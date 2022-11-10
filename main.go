package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	handler_enter := http.HandlerFunc(handleRequest_enter)
	http.Handle("/", handler_enter)
	handler_health := http.HandlerFunc(handleRequest_health)
	http.Handle("/health", handler_health)
	http.ListenAndServe(":8888", nil)
}

func handleRequest_health(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["status"] = "ok"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func handleRequest_enter(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := ioutil.ReadFile("Hello.jpg")
	if err != nil {
		panic(err)
	}
	w.Write(fileBytes)
	return
}
