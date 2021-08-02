package code

import (
	"context"
	"time"
)

var (
	fakePaymentCode = "abcd1234"
	fakeName        = "Local Test"
)

func randomPaymentCode() PaymentCode {
	return PaymentCode{
		ID:             1,
		PaymentCode:    fakePaymentCode,
		Name:           fakeName,
		Status:         "ACTIVE",
		CreatedAt:      time.Now().UTC().Truncate(1 * time.Second),
		UpdatedAt:      time.Now().UTC().Truncate(1 * time.Second),
		ExpirationDate: time.Now().UTC().AddDate(50, 0, 0).Format(time.RFC3339),
	}
}

type FakeService struct {
	findPaymentCodeByIdFn   func(id string) (PaymentCode, error)
	findPaymentCodeByCodeFn func(code string) (PaymentCode, error)
	generatePaymentCodeFn   func(reqBody reqBodyPaymentCode) (PaymentCode, error)
}

func (f FakeService) FindPaymentCodeById(id string) (PaymentCode, error) {
	return f.findPaymentCodeByIdFn(id)
}

func (f FakeService) FindPaymentCodeByCode(code string) (PaymentCode, error) {
	return f.findPaymentCodeByCodeFn(code)
}

func (f FakeService) GeneratePaymentCode(reqBody reqBodyPaymentCode) (PaymentCode, error) {
	return f.generatePaymentCodeFn(reqBody)
}

type fakeRepository struct {
	findPaymentCodeByIdFn func(ctx context.Context, id int) (PaymentCode, error)
	generatePaymentCodeFn func(ctx context.Context, request PaymentCode) (PaymentCode, error)
}

func (f fakeRepository) FindPaymentCodeById(ctx context.Context, id int) (PaymentCode, error) {
	return f.findPaymentCodeByIdFn(ctx, id)
}

func (f fakeRepository) GeneratePaymentCode(ctx context.Context, request PaymentCode) (PaymentCode, error) {
	return f.generatePaymentCodeFn(ctx, request)
}
