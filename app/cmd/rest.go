package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func StartRest() {
	controller()
}

func controller() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/hello-world", helloWorld)

	fmt.Print("Starting server on port 3000\n")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.URL.Path != "health" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	type response struct {
		Status string `json:"status"`
	}

	data := response{"healthy"}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.URL.Path != "hello-world" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	type response struct {
		Message string `json:"message"`
	}

	data := response{"hello world"}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	if status == http.StatusNotFound {
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error_code": "Page not found"}`))
	}
}
