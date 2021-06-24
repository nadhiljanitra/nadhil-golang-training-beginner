package expirecode

import (
	"database/sql"
	"fmt"
)

type repository interface {
	GetExpiredPaymentCode() error
}

type sqlRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository {
	return sqlRepository{
		db: db,
	}
}

func (s sqlRepository) GetExpiredPaymentCode() error {
	fmt.Printf("Succesfully reaching repository")
	return nil
}
