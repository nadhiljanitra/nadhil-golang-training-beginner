package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/app/config"
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/healthcheck"
	code "github.com/nadhiljanitra/nadhil-golang-training-beginner/paymentcode"
)

func InitRest() {
	// init postgres database
	db, err := config.InitPostgres()
	if err != nil {
		panic(err)
	}

	paymentRepository := code.NewSQLRepository(db)
	paymentService := code.NewService(paymentRepository)

	restHandler(paymentService)
}

func restHandler(paymentSvc code.Service) {
	// healthcheck Controller
	healthcheck.RegisterCtrl()

	// paymentCode Controller
	code.RegisterCtrl(paymentSvc)

	//TODO update the logger in here
	fmt.Print("Starting server on port 3000\n")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
