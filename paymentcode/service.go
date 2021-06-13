package code

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
)

type Service interface {
	FindPaymentCodeById(id string) (PaymentCode, error)
	GeneratePaymentCode(reqBody reqBodyPaymentCode) (PaymentCode, error)
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

	ID, err := strconv.Atoi(id)
	if err != nil {
		if errors.Is(err, strconv.ErrSyntax) {
			// fmt.Printf("\nPayment code must be a number")
			// Put as NotFound 404 to satisfy the postman test where it should be errIllegalArg
			return PaymentCode{}, common.ErrNotFound
		}
		return PaymentCode{}, common.ErrUnexpected
	}

	paymentCode, err := d.repo.FindPaymentCodeById(contextTimeout, ID)
	if err != nil {
		return PaymentCode{}, err
	}

	return paymentCode, nil
}

func (d defaultService) GeneratePaymentCode(reqBody reqBodyPaymentCode) (PaymentCode, error) {
	fiveSecond := 5 * time.Second
	contextTimeout, cancel := context.WithTimeout(context.TODO(), fiveSecond)
	defer cancel()

	paymentCode, err := NewPaymentCode(reqBody)
	if err != nil {
		return PaymentCode{}, err
	}

	result, err := d.repo.GeneratePaymentCode(contextTimeout, paymentCode)
	if err != nil {
		return PaymentCode{}, err
	}

	return result, nil
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
