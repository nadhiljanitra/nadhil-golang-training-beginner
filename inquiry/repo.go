package inquiry

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
)

type repository interface {
	CreateInquiry(ctx context.Context, params CreateInquiryParam) error
	GetInquiry(ctx context.Context, transactionID string) error
}

func NewSQLRepository(db *sql.DB) repository {
	return sqlRepository{
		DB: db,
	}
}

type sqlRepository struct {
	DB *sql.DB
}

type CreateInquiryParam struct {
	ID            string
	TransactionID string
	PaymentCode   string
}

func (s sqlRepository) CreateInquiry(ctx context.Context, i CreateInquiryParam) error {
	_, err := s.DB.Exec(`INSERT into inquiries (id, transaction_id, payment_code) values ($1,$2,$3)`, i.ID, i.TransactionID, i.PaymentCode)
	if err, ok := err.(*pq.Error); ok {
		// https://www.postgresql.org/docs/9.3/errcodes-appendix.html
		if err.Code.Name() == "unique_violation" {
			return common.ErrDuplicateInquiry
		}

		return common.ErrUnexpected
	}

	return nil
}

func (s sqlRepository) GetInquiry(ctx context.Context, transactionID string) error {
	return nil
}
