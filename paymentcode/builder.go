package code

import (
	"time"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
)

type PaymentCode struct {
	// ID type put as int to satisfied the postman test
	ID int `json:"id"`

	PaymentCode    string    `json:"payment_code"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	ExpirationDate string    `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func NewPaymentCode(bodyRequest reqBodyPaymentCode) (PaymentCode, error) {
	if bodyRequest.Name == nil || bodyRequest.PaymentCode == nil {
		return PaymentCode{}, common.ErrIllegalArg
	}

	paymentCode := PaymentCode{
		PaymentCode:    *bodyRequest.PaymentCode,
		Name:           *bodyRequest.Name,
		Status:         "ACTIVE",
		CreatedAt:      time.Now().UTC().Truncate(1 * time.Second),
		UpdatedAt:      time.Now().UTC().Truncate(1 * time.Second),
		ExpirationDate: time.Now().UTC().AddDate(50, 0, 0).Format(time.RFC3339),
	}

	return paymentCode, nil
}
