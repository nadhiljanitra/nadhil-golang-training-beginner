package inquiry

import (
	"encoding/json"
	"net/http"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
	code "github.com/nadhiljanitra/nadhil-golang-training-beginner/paymentcode"
)

type endpoint struct {
	codeService    code.Service
	inquiryService Service
}

type RequestBody struct {
	TransactionID string `json:"transaction_id"`
	PaymentCode   string `json:"payment_code"`
}

type ResponseBody struct {
	PaymentCode string `json:"payment_code"`
	Amount      int    `json:"amount"`
	Name        string `json:"name"`
	Status      string `json:"status"`
}

func RegisterCtrl(paymentCodeService code.Service, inquiryService Service) {
	ctrl := endpoint{
		codeService:    paymentCodeService,
		inquiryService: inquiryService,
	}

	http.HandleFunc("/inquiry", ctrl.CreateInquiry)
}

func (e endpoint) CreateInquiry(w http.ResponseWriter, r *http.Request) {
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

	paymentCode, err := e.codeService.FindPaymentCodeByCode(reqBody.PaymentCode)
	if err != nil {
		common.ErrorHandler(err, w, r)
		return
	}

	if paymentCode.Status != code.Active {
		return
	}

	if err = e.inquiryService.CreateInquiry(reqBody.TransactionID, reqBody.PaymentCode); err != nil {
		common.ErrorHandler(err, w, r)
		return
	}

	resBody := ResponseBody{
		PaymentCode: paymentCode.PaymentCode,
		// TODO add amount on payment code
		Amount: 0,
		Name:   paymentCode.Name,
		Status: paymentCode.Status,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resBody)
	return
}
