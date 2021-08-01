package code

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
)

func TestGetPaymentCodeByID(t *testing.T) {
	//preparation
	paymentCodeHandler := paymentEndpoint{
		codeService: FakeService{
			findPaymentCodeByIdFn: func(id string) (PaymentCode, error) {
				return randomPaymentCode(), nil
			},
		},
	}

	req, err := http.NewRequest("GET", "/payment-codes/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(paymentCodeHandler.getPaymentCodeById)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	responseJson := PaymentCode{}
	err = json.Unmarshal(rr.Body.Bytes(), &responseJson)
	if err != nil {
		t.Errorf("Failed on unmarshalling response")
	}

	expected := randomPaymentCode()
	if responseJson != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseJson, expected)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v", ctype, "application/json")
	}
}

func TestFailedPaymentCodeNotFound(t *testing.T) {
	//preparation
	paymentCodeHandler := paymentEndpoint{
		codeService: FakeService{
			findPaymentCodeByIdFn: func(id string) (PaymentCode, error) {
				return PaymentCode{}, common.ErrNotFound
			},
		},
	}

	req, err := http.NewRequest("GET", "/payment-codes/1456", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(paymentCodeHandler.getPaymentCodeById)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	responseJson := common.HTTPErrorResponse{}
	err = json.Unmarshal(rr.Body.Bytes(), &responseJson)
	if err != nil {
		t.Errorf("Failed on unmarshalling response")
	}

	expected := common.HTTPErrorResponse{
		ErrorCode: "DATA_NOT_FOUND_ERROR",
		Message:   common.ErrNotFound.Error(),
	}

	if responseJson != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseJson, expected)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v", ctype, "application/json")
	}
}

func TestGeneratePaymentCode(t *testing.T) {
	//preparation
	paymentCodeHandler := paymentEndpoint{
		codeService: FakeService{
			generatePaymentCodeFn: func(reqBody reqBodyPaymentCode) (PaymentCode, error) {
				return randomPaymentCode(), nil
			},
		},
	}

	requestBody := reqBodyPaymentCode{
		Name:        &fakeName,
		PaymentCode: &fakePaymentCode,
	}
	reqBodyBytes, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "/payment-codes", bytes.NewReader(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(paymentCodeHandler.generatePaymentCode)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	responseJson := PaymentCode{}
	err = json.Unmarshal(rr.Body.Bytes(), &responseJson)
	if err != nil {
		t.Errorf("Failed on unmarshalling response")
	}

	expected := randomPaymentCode()
	if responseJson != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", responseJson, expected)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v", ctype, "application/json")
	}
}
