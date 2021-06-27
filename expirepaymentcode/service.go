package expirecode

import (
	"context"
	"time"
)

type Service interface {
	ExpiringPaymentCode() (int64, error)
}

type defaultService struct {
	repo repository
}

func NewService(repo repository) Service {
	return defaultService{
		repo: repo,
	}
}

func (s defaultService) ExpiringPaymentCode() (int64, error) {
	fiveSecond := 5 * time.Second
	contextTimeout, cancel := context.WithTimeout(context.TODO(), fiveSecond)
	defer cancel()

	now := time.Now()
	updated, err := s.repo.ExpiringPaymentCode(contextTimeout, now)
	if err != nil {
		return 0, err
	}
	return updated, nil
}
