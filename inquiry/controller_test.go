package inquiry

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	code "github.com/nadhiljanitra/nadhil-golang-training-beginner/paymentcode"
)

func TestCreateInquiry(t *testing.T) {
	//preparation
	inquiryHandler := endpoint{
		inquiryService: FakeService{
			createInquiryfn: func(transactionID, paymentCode string) error {
				return nil
			}},
		codeService: code.FakeService{
			FindPaymentCodeByCodeFn: func(codes string) (code.PaymentCode, error) {
				return code.RandomPaymentCode(), nil
			},
		},
	}

	requestBody := RequestBody{
		TransactionID: "transaction123",
		PaymentCode:   "test1234",
	}
	reqBodyBytes, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "/inquiry", bytes.NewReader(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(inquiryHandler.CreateInquiry)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	responseJson := ResponseBody{}
	err = json.Unmarshal(rr.Body.Bytes(), &responseJson)
	if err != nil {
		t.Errorf("Failed on unmarshalling response")
	}

	expected := ResponseBody{
		PaymentCode: "abcd1234",
		Amount:      10,
		Name:        "Local Test",
		Status:      "ACTIVE",
	}
	if responseJson != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseJson, expected)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v", ctype, "application/json")
	}
}
