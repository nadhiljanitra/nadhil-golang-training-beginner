package code

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
)

type repository interface {
	FindPaymentCodeById(ctx context.Context, id int) (PaymentCode, error)
	GeneratePaymentCode(ctx context.Context, request PaymentCode) (PaymentCode, error)
}

func NewSQLRepository(db *sql.DB) repository {
	return sqlRepository{
		DB: db,
	}
}

type sqlRepository struct {
	DB *sql.DB
}

func (s sqlRepository) FindPaymentCodeById(ctx context.Context, id int) (PaymentCode, error) {
	var ID int
	var paymentCode string
	var name string
	var status string
	var expirationDate string
	var createdAt time.Time
	var updatedAt time.Time

	row := s.DB.QueryRowContext(ctx, `
	SELECT 
	id, 
	payment_code, 
	name, 
	status, 
	expiration_date, 
	created_at, 
	updated_at 
	FROM payment_codes 
	WHERE id=$1`,
		id)

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
			return PaymentCode{}, common.ErrNotFound
		}

		return PaymentCode{}, err
	}

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

func (s sqlRepository) GeneratePaymentCode(ctx context.Context, r PaymentCode) (PaymentCode, error) {
	var ID int

	result := s.DB.QueryRowContext(ctx, `
	INSERT into payment_codes 
	(payment_code, name, status, expiration_date, created_at, updated_at) 
	values ($1,$2,$3,$4,$5,$6) 
	RETURNING id`,
		r.PaymentCode, r.Name, r.Status, r.ExpirationDate, r.CreatedAt, r.UpdatedAt)
	err := result.Scan(&ID)
	if err != nil {
		return PaymentCode{}, err
	}

	paymentCode := PaymentCode{
		ID:             ID,
		PaymentCode:    r.PaymentCode,
		Name:           r.Name,
		Status:         r.Status,
		ExpirationDate: r.ExpirationDate,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
	}

	return paymentCode, nil
}
