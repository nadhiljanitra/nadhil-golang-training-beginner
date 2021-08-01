package code

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
)

type paymentEndpoint struct {
	codeService Service
}

type reqBodyPaymentCode struct {
	Name        *string `json:"name,omitempty"`
	PaymentCode *string `json:"payment_code,omitempty"`
}

func RegisterCtrl(paymentSvc Service) {
	ctrl := paymentEndpoint{
		codeService: paymentSvc,
	}

	http.HandleFunc("/", ctrl.router)
}

func (e paymentEndpoint) router(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		e.getPaymentCodeById(w, r)
	case http.MethodPost:
		e.generatePaymentCode(w, r)
	default:
		http.NotFound(w, r)
	}

}

func (e paymentEndpoint) getPaymentCodeById(w http.ResponseWriter, r *http.Request) {
	if !validateBaseURL(r.URL.Path) {
		http.NotFound(w, r)
		return
	}

	paymentCodeID := getPaymentCodeIDFromUrl(r.URL.Path)
	paymentCode, err := e.codeService.FindPaymentCodeById(paymentCodeID)
	if err != nil {
		common.ErrorHandler(err, w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentCode)
	return
}

func (e paymentEndpoint) generatePaymentCode(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/payment-codes" {
		http.NotFound(w, r)
		return
	}

	var reqBody reqBodyPaymentCode

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		common.ErrorHandler(common.ErrIllegalArg, w, r)
		return
	}

	paymentCode, err := e.codeService.GeneratePaymentCode(reqBody)
	if err != nil {
		common.ErrorHandler(err, w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(paymentCode)
	return
}

func getPaymentCodeIDFromUrl(url string) string {
	var baseUrl = "/payment-codes/"
	paymentCode := strings.TrimPrefix(url, baseUrl)

	return paymentCode
}

func validateBaseURL(url string) bool {
	var baseUrl = "/payment-codes"
	return strings.HasPrefix(url, baseUrl)
}
