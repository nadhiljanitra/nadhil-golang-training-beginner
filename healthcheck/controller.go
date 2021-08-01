package healthcheck

import (
	"encoding/json"
	"net/http"
)

func RegisterCtrl() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/hello-world", helloWorld)
}

type healthCheckResponse struct {
	Status string `json:"status"`
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	data := healthCheckResponse{"healthy"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	data := helloWorldResponse{"hello world"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
