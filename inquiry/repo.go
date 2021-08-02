package inquiry

import (
	"context"
	"database/sql"
)

type repository interface {
	CreateInquiry(ctx context.Context, transactionID string, paymentCode string) error
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

func (s sqlRepository) CreateInquiry(ctx context.Context, transactionID string, paymentCode string) error {
	return nil
}

func (s sqlRepository) GetInquiry(ctx context.Context, transactionID string) error {
	return nil
}
