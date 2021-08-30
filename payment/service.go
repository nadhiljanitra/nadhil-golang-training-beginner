package payment

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/payment/publisher"
)

type Service interface {
	CreatePayment(payment Payment) error
}

type defaultService struct {
	repo      repository
	publisher publisher.Publisher
}

func NewService(repo repository, publisher publisher.Publisher) Service {
	return defaultService{
		repo:      repo,
		publisher: publisher,
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

	message, _ := json.Marshal(payment)
	if err := s.publisher.Publish(ctx, message); err != nil {
		return err
	}

	return nil
}
