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
