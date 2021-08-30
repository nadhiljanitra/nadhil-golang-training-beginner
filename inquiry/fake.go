package inquiry

import (
	"context"

	"github.com/google/uuid"
)

type fakeRepository struct {
	createInquiryfn             func(ctx context.Context, params CreateInquiryParam) error
	getInquiryByTransactionIDfn func(ctx context.Context, transactionID string) (Inquiry, error)
}

func (f fakeRepository) CreateInquiry(ctx context.Context, params CreateInquiryParam) error {
	return f.createInquiryfn(ctx, params)
}

func (f fakeRepository) GetInquiryByTransactionID(ctx context.Context, transactionID string) (Inquiry, error) {
	return f.getInquiryByTransactionIDfn(ctx, transactionID)
}

type FakeService struct {
	createInquiryfn             func(transactionID string, paymentCode string) error
	getInquiryByTransactionIDfn func(transactionID string) (Inquiry, error)
}

func (f FakeService) CreateInquiry(transactionID string, paymentCode string) error {
	return f.createInquiryfn(transactionID, paymentCode)
}

func (f FakeService) GetInquiryByTransactionID(transactionID string) (Inquiry, error) {
	return f.getInquiryByTransactionIDfn(transactionID)
}

func randomInquiry() Inquiry {
	return Inquiry{
		ID:            uuid.New().String(),
		TransactionID: "transaction123",
		PaymentCode:   "abcd1234",
	}
}
