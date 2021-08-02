package inquiry

import (
	"context"
	"time"
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

	err := s.repo.CreateInquiry(ctx, transactionID, paymentCode)
	if err != nil {
		return err
	}

	return nil
}
