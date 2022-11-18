package handlers

import (
	"encoding/json"
	"io"
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
	jsonResp, _ := json.Marshal(response)

	responseError := map[string]string{
		"status": "error",
		"error":  "Bad Request. Message must have body",
	}
	jsonRespErr, _ := json.Marshal(responseError)

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res := Message{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		w.Write(jsonRespErr)
		return
	}
	if res.Title == "" {
		w.Write(jsonRespErr) // Если заголовок сообщения пустой
		return
	}
	w.Write(jsonResp)
}
