package payment

import (
	"encoding/json"
	"net/http"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/inquiry"
	code "github.com/nadhiljanitra/nadhil-golang-training-beginner/paymentcode"
)

type endpoint struct {
	codeService    code.Service
	inquiryService inquiry.Service
	paymentService Service
}

type RequestBody struct {
	TransactionID string `json:"transaction_id"`
	PaymentCode   string `json:"payment_code"`
	Amount        int    `json:"amount"`
	Name          string `json:"name"`
}

type ResponseBody struct {
	TransactionID string `json:"transaction_id"`
	Name          string `json:"name"`
	Amount        int    `json:"amount"`
	Status        string `json:"status"`
}

func RegisterCtrl(paymentCodeService code.Service, inquiryService inquiry.Service, paymentService Service) {
	ctrl := endpoint{
		codeService:    paymentCodeService,
		inquiryService: inquiryService,
		paymentService: paymentService,
	}

	http.HandleFunc("/payment", ctrl.CreatePayment)
}

func (e endpoint) CreatePayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	var reqBody RequestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		common.ErrorHandler(common.ErrIllegalArg, w, r)
		return
	}

	inquiryData, err := e.inquiryService.GetInquiryByTransactionID(reqBody.TransactionID)
	if err != nil {
		common.ErrorHandler(err, w, r)
		return
	}

	payment := Payment{
		TransactionID: reqBody.TransactionID,
		PaymentCode:   reqBody.PaymentCode,
		Name:          reqBody.Name,
		Amount:        reqBody.Amount,
	}

	if err := e.paymentService.CreatePayment(payment); err != nil {
		common.ErrorHandler(err, w, r)
		return
	}

	resBody := ResponseBody{
		TransactionID: inquiryData.TransactionID,
		Amount:        reqBody.Amount,
		Name:          reqBody.Name,
		Status:        "SUCCESS",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resBody)
	return
}
