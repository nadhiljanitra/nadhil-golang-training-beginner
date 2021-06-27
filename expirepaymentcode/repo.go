package expirecode

import (
	"context"
	"database/sql"
	"time"

	code "github.com/nadhiljanitra/nadhil-golang-training-beginner/paymentcode"
)

type repository interface {
	ExpiringPaymentCode(ctx context.Context, now time.Time) (int64, error)
}

type sqlRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) repository {
	return sqlRepository{
		DB: db,
	}
}

func (s sqlRepository) ExpiringPaymentCode(ctx context.Context, now time.Time) (int64, error) {
	query := `UPDATE payment_codes 
	SET status = $1, updated_at = $3 
	WHERE status = $2 
	AND expiration_date < $3`

	result, err := s.DB.ExecContext(ctx, query, code.Expired, code.Active, now)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}
