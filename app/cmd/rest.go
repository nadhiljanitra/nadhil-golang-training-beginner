package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/app/config"
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/healthcheck"
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/inquiry"
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

	inquiryRepository := inquiry.NewSQLRepository(db)
	inquiryService := inquiry.NewService(inquiryRepository)

	restHandler(paymentService, inquiryService)
}

func restHandler(paymentSvc code.Service, inquirySvc inquiry.Service) {
	// healthcheck Controller
	healthcheck.RegisterCtrl()

	// paymentCode Controller
	code.RegisterCtrl(paymentSvc)

	// inquiry Controller
	inquiry.RegisterCtrl(paymentSvc, inquirySvc)

	//TODO update the logger in here
	fmt.Print("Starting server on port 3000\n")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
