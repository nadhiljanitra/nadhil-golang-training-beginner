package inquiry

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateInquiry(transactionID string, paymentCode string) error
}

type defaultService struct {
	repo repository
}

func NewService(repo repository) Service {
	return defaultService{
		repo: repo,
	}
}

func (s defaultService) CreateInquiry(transactionID string, paymentCode string) error {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	ID := uuid.New().String()

	params := CreateInquiryParam{
		ID:            ID,
		TransactionID: transactionID,
		PaymentCode:   paymentCode,
	}

	if err := s.repo.CreateInquiry(ctx, params); err != nil {
		return err
	}

	return nil
}
