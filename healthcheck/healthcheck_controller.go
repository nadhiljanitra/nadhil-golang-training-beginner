package healthcheck

import (
	"encoding/json"
	"net/http"
)

func RegisterCtrl() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/hello-world", helloWorld)
}

func health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	type response struct {
		Status string `json:"status"`
	}

	data := response{"healthy"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	type response struct {
		Message string `json:"message"`
	}

	data := response{"hello world"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
