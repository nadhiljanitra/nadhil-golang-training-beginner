package payment

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreatePayment(payment Payment) error
}

type defaultService struct {
	repo repository
}

func NewService(repo repository) Service {
	return defaultService{
		repo: repo,
	}
}

type Payment struct {
	TransactionID string
	PaymentCode   string
	Name          string
	Amount        int
}

func (s defaultService) CreatePayment(payment Payment) error {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	ID := uuid.New().String()

	params := CreatePaymentParam{
		ID:          ID,
		PaymentData: payment,
	}

	if err := s.repo.CreatePayment(ctx, params); err != nil {
		return err
	}

	return nil
}
