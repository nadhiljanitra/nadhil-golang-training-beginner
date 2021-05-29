package code

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type PaymentCode struct {
	ID             string    `json:"id"`
	PaymentCode    string    `json:"payment_code"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	ExpirationDate string    `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Service interface {
	FindPaymentCodeById(id string) (PaymentCode, error)
}

type defaultService struct {
	repo repository
}

func NewService(repo repository) Service {
	return defaultService{
		repo: repo,
	}
}

func (d defaultService) FindPaymentCodeById(id string) (PaymentCode, error) {
	fiveSecond := 5 * time.Second

	contextTimeout, cancel := context.WithTimeout(context.TODO(), fiveSecond)
	defer cancel()

	paymentCode, err := d.repo.FindPaymentCodeById(contextTimeout, id)
	if err != nil {
		return PaymentCode{}, err
	}

	fmt.Println("\n PAYMENT CODE =====<> ", paymentCode)

	return PaymentCode{}, nil
}

func getPaymentCodeIDFromUrl(url string) string {
	baseUrl := "/payment-codes/"
	paymentCode := strings.TrimPrefix(url, baseUrl)

	return paymentCode
}
