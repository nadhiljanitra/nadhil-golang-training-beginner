package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/healthcheck"
)

func InitRest() {
	controller()
}

func controller() {
	healthcheck.RegisterCtrl()

	fmt.Print("Starting server on port 3000\n")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
