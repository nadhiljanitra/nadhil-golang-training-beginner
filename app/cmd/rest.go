package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/app/config"
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/healthcheck"
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/inquiry"
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/payment"
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/payment/publisher"
	code "github.com/nadhiljanitra/nadhil-golang-training-beginner/paymentcode"
)

func InitRest() {
	// init postgres database
	db, err := config.InitPostgres()
	if err != nil {
		panic(err)
	}

	queueURL := os.Getenv("SQS_QUEUE_URL")

	codeRepository := code.NewSQLRepository(db)
	codeService := code.NewService(codeRepository)

	inquiryRepository := inquiry.NewSQLRepository(db)
	inquiryService := inquiry.NewService(inquiryRepository)

	sqsClient := config.InitSQSClient()
	publisher := publisher.NewSQSPublisher(sqsClient, queueURL)

	paymentRepository := payment.NewSQLRepository(db)
	paymentService := payment.NewService(paymentRepository, publisher)

	restHandler(codeService, inquiryService, paymentService)
}

func restHandler(codeService code.Service, inquirySvc inquiry.Service, paymentSvc payment.Service) {
	// healthcheck Controller
	healthcheck.RegisterCtrl()

	// paymentCode Controller
	code.RegisterCtrl(codeService)

	// inquiry Controller
	inquiry.RegisterCtrl(codeService, inquirySvc)

	// payment Controller
	payment.RegisterCtrl(codeService, inquirySvc, paymentSvc)

	//TODO update the logger in here
	fmt.Print("Starting server on port 3000\n")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
