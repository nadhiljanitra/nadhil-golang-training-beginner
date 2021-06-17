package code

import (
	"context"
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/internal/postgres"
	"github.com/stretchr/testify/suite"
)

type repoIntegrationSuite struct {
	postgres.Suite
}

func TestRepoIntegrationSuite(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	// setup migration file directory, may be different based on the migration file directory
	migrationsDir := path.Join(filepath.Dir(b), "..", "..", "nadhil-golang-training-beginner", "internal", "postgres", "migrations")
	repoSuite := postgres.NewDefaultSuite(migrationsDir)

	slipSuite := &repoIntegrationSuite{*repoSuite}

	suite.Run(t, slipSuite)
}

func (s repoIntegrationSuite) TestInsertPaymentCode() {
	fiveSecond := 5 * time.Second
	contextTimeout, cancel := context.WithTimeout(context.TODO(), fiveSecond)
	defer cancel()

	var name string = "test"
	var pCode string = "test1234"

	req := reqBodyPaymentCode{
		Name:        &name,
		PaymentCode: &pCode,
	}

	paymentCode, err := NewPaymentCode(req)
	s.Require().Nil(err)

	repo := sqlRepository{
		DB: s.DB(),
	}
	result, err := repo.GeneratePaymentCode(contextTimeout, paymentCode)
	s.Require().Nil(err)

	ID := result.ID
	query := fmt.Sprintf("select count(*) from payment_codes where ID='%d'", ID)
	row := s.DB().QueryRow(query)
	var count int
	err = row.Scan(&count)
	s.Require().Nil(err)
	s.Require().Equal(1, count)
}

func (s repoIntegrationSuite) insertPaymentCode(ctx context.Context, name string, code string) PaymentCode {
	req := reqBodyPaymentCode{
		Name:        &name,
		PaymentCode: &code,
	}

	paymentCode, err := NewPaymentCode(req)
	s.Require().Nil(err)

	repo := sqlRepository{
		DB: s.DB(),
	}
	result, err := repo.GeneratePaymentCode(ctx, paymentCode)
	s.Require().Nil(err)

	return result
}

func (s repoIntegrationSuite) TestFindPaymentCodeById() {
	fiveSecond := 5 * time.Second
	contextTimeout, cancel := context.WithTimeout(context.TODO(), fiveSecond)
	defer cancel()

	generatedPaymentCode := s.insertPaymentCode(contextTimeout, "testing", "testing1234")

	repo := sqlRepository{
		DB: s.DB(),
	}
	result, err := repo.FindPaymentCodeById(contextTimeout, generatedPaymentCode.ID)
	s.Require().Nil(err)
	s.Require().Equal(generatedPaymentCode, result)
}

func (s repoIntegrationSuite) TestFindPaymentCodeByIdNotFound() {
	fiveSecond := 5 * time.Second
	contextTimeout, cancel := context.WithTimeout(context.TODO(), fiveSecond)
	defer cancel()

	repo := sqlRepository{
		DB: s.DB(),
	}
	result, err := repo.FindPaymentCodeById(contextTimeout, 999999999)
	s.Require().Equal(result, PaymentCode{})
	s.Require().NotNil(err)
	s.Require().Equal(err, common.ErrNotFound)
}
