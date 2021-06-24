package expirecode

import "fmt"

type Service interface {
	FindExpiredPaymentCode() error
}

type defaultService struct {
	repo repository
}

func NewService(repo repository) Service {
	return defaultService{
		repo: repo,
	}
}

func (s defaultService) FindExpiredPaymentCode() error {
	fmt.Printf("Started on service level\n")
	s.repo.GetExpiredPaymentCode()
	return nil
}
