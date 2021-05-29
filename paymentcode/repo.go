package code

import (
	"context"
	"database/sql"
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

	rows, err := s.DB.QueryContext(ctx, "SELECT id, payment_code, name, status, expiration_date FROM payment_codes WHERE id=%?", id)
	if err != nil {
		//TODO update the logger here
		fmt.Println("Error on pq, ", err)
		return PaymentCode{}, fmt.Errorf("%w", common.ErrUnexpected)
	}
	defer rows.Close()

	err = rows.Scan(
		&ID,
		&paymentCode,
		&name,
		&status,
		&expirationDate,
		&createdAt,
		&updatedAt,
	)

	result := PaymentCode{
		ID:             ID,
		PaymentCode:    paymentCode,
		Name:           name,
		Status:         status,
		ExpirationDate: expirationDate,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}

	return result, nil
}
