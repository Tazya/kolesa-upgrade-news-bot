package delievery

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	if err := h.tmpl.ExecuteTemplate(w, "home.html", nil); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"status": "ok",
	}
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	w.Write(jsonResp)
	return
}
