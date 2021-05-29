package code

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
)

type paymentEndpoint struct {
	codeService Service
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
	default:
		http.NotFound(w, r)
	}

}

func (e paymentEndpoint) getPaymentCodeById(w http.ResponseWriter, r *http.Request) {
	paymentCodeID := getPaymentCodeIDFromUrl(r.URL.Path)

	paymentCode, err := e.codeService.FindPaymentCodeById(paymentCodeID)
	if err != nil {
		common.ErrorHandler(err, w, r)
		return
	}

	fmt.Println("\npaymentCode ", paymentCode)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentCode)
	return
}
