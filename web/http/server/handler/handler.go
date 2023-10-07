package handler

import (
	"fmt"
	"net/http"
)

func PostDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post request are allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	data := r.FormValue("data")
	response := fmt.Sprintf("Received data: %s", data)
	fmt.Fprintf(w, response)
}
