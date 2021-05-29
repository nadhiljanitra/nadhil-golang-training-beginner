package code

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
)

type repository interface {
	FindPaymentCodeById(ctx context.Context, id string) (PaymentCode, error)
}

func NewSQLRepository(db *sql.DB) repository {
	return sqlRepository{
		DB: db,
	}
}

type sqlRepository struct {
	DB *sql.DB
}

func (s sqlRepository) FindPaymentCodeById(ctx context.Context, id string) (PaymentCode, error) {
	var ID string
	var paymentCode string
	var name string
	var status string
	var expirationDate string
	var createdAt time.Time
	var updatedAt time.Time

	result := PaymentCode{
		ID:             ID,
		PaymentCode:    paymentCode,
		Name:           name,
		Status:         status,
		ExpirationDate: expirationDate,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}

	row := s.DB.QueryRowContext(ctx, "SELECT id, payment_code, name, status, expiration_date FROM payment_codes WHERE id=$1", id)

	err := row.Scan(
		&ID,
		&paymentCode,
		&name,
		&status,
		&expirationDate,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, fmt.Errorf("%w", common.ErrNotFound)
		}
		return result, err
	}

	return result, nil
}
