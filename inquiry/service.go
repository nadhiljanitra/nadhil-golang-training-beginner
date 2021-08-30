package inquiry

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateInquiry(transactionID string, paymentCode string) error
	GetInquiryByTransactionID(transactionID string) (Inquiry, error)
}

type defaultService struct {
	repo repository
}

func NewService(repo repository) Service {
	return defaultService{
		repo: repo,
	}
}

type Inquiry struct {
	ID            string
	TransactionID string
	PaymentCode   string
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

func (s defaultService) GetInquiryByTransactionID(transactionID string) (Inquiry, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	inquiry, err := s.repo.GetInquiryByTransactionID(ctx, transactionID)
	if err != nil {
		return Inquiry{}, err
	}

	return inquiry, nil
}
