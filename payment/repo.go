package payment

import (
	"context"
	"database/sql"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
)

type repository interface {
	CreatePayment(ctx context.Context, params CreatePaymentParam) error
}

func NewSQLRepository(db *sql.DB) repository {
	return sqlRepository{
		DB: db,
	}
}

type sqlRepository struct {
	DB *sql.DB
}

type CreatePaymentParam struct {
	ID          string
	PaymentData Payment
}

func (s sqlRepository) CreatePayment(ctx context.Context, p CreatePaymentParam) error {
	_, err := s.DB.Exec(`INSERT into payments (id, transaction_id, payment_code, name, amount) values ($1,$2,$3,$4,$5)`,
		p.ID,
		p.PaymentData.TransactionID, p.PaymentData.PaymentCode, p.PaymentData.Name, p.PaymentData.Amount)
	if err != nil {
		return common.ErrUnexpected
	}

	return nil
}
