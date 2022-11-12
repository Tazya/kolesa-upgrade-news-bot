package handlers

import (
	"net/http"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fileBytes, err := os.ReadFile("Hello.jpg")
	if err != nil {
		panic(err)
	}
	w.Write(fileBytes)
}
