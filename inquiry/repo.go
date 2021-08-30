package inquiry

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
)

type repository interface {
	CreateInquiry(ctx context.Context, params CreateInquiryParam) error
	GetInquiryByTransactionID(ctx context.Context, transactionID string) (Inquiry, error)
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

func (s sqlRepository) GetInquiryByTransactionID(ctx context.Context, transactionID string) (Inquiry, error) {
	var ID string
	var trxID string
	var paymentCode string

	row := s.DB.QueryRowContext(ctx, `
	SELECT 
	id, 
	transaction_id, 
	payment_code
	FROM inquiries 
	WHERE transaction_id=$1`,
		transactionID)

	err := row.Scan(
		&ID,
		&trxID,
		&paymentCode,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Inquiry{}, common.ErrNotFound
		}

		return Inquiry{}, common.ErrUnexpected
	}

	result := Inquiry{
		ID:            ID,
		PaymentCode:   paymentCode,
		TransactionID: trxID,
	}

	return result, nil
}
