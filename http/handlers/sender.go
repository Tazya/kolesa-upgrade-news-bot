package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Sender(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "ok",
	}
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	responseError := map[string]string{
		"status": "error",
		"error":  "Bad Request. Message must have body",
	}
	jsonRespErr, err := json.Marshal(responseError)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res := Message{}
	err1 := json.Unmarshal(body, &res)
	if err1 != nil {
		log.Printf("Ошибка при декодировании сообщений %v", err)
		return
	}
	if res.Title == "" {
		w.Write(jsonRespErr) // Если заголовок сообщения пустой
		return
	}
	w.Write(jsonResp)
}
