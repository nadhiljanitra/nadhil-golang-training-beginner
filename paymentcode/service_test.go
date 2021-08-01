package code

import (
	"context"
	"time"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Create a testing using ginkgo <> gomega package
var _ = Describe("Service", func() {
	Context("Find payment code by ID", func() {
		var createdPaymentCode PaymentCode

		BeforeEach(func() {
			createdPaymentCode = PaymentCode{
				ID:             1,
				PaymentCode:    "test1234",
				Name:           "TEST",
				Status:         "ACTIVE",
				CreatedAt:      time.Now().UTC().Truncate(1 * time.Second),
				UpdatedAt:      time.Now().UTC().Truncate(1 * time.Second),
				ExpirationDate: time.Now().UTC().AddDate(50, 0, 0).Format(time.RFC3339),
			}
		})

		It("Should find payment code", func() {
			repo := fakeRepository{
				findPaymentCodeByIdFn: func(ctx context.Context, id int) (PaymentCode, error) {
					return createdPaymentCode, nil
				},
			}
			svc := NewService(repo)

			actual, err := svc.FindPaymentCodeById("1")
			Expect(err).To(Not(HaveOccurred()))
			Expect(actual).To(Equal(createdPaymentCode))
		})

		It("Should failed for payment code not a number", func() {
			repo := fakeRepository{}
			svc := NewService(repo)

			_, err := svc.FindPaymentCodeById("abcd")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(common.ErrNotFound))
		})

		It("Should failed for payment code not found", func() {
			repo := fakeRepository{
				findPaymentCodeByIdFn: func(ctx context.Context, id int) (PaymentCode, error) {
					return PaymentCode{}, common.ErrNotFound
				},
			}

			svc := NewService(repo)
			_, err := svc.FindPaymentCodeById("123")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(common.ErrNotFound))
		})
	})

	Context("Generate payment code", func() {
		var createdPaymentCode PaymentCode

		BeforeEach(func() {
			createdPaymentCode = PaymentCode{
				ID:             1,
				PaymentCode:    "test1234",
				Name:           "TEST",
				Status:         "ACTIVE",
				CreatedAt:      time.Now().UTC().Truncate(1 * time.Second),
				UpdatedAt:      time.Now().UTC().Truncate(1 * time.Second),
				ExpirationDate: time.Now().UTC().AddDate(50, 0, 0).Format(time.RFC3339),
			}
		})

		It("Should generate payment code", func() {
			var name string = "TEST"
			var code string = "test1234"
			reqBody := reqBodyPaymentCode{
				Name:        &name,
				PaymentCode: &code,
			}

			repo := fakeRepository{
				generatePaymentCodeFn: func(ctx context.Context, request PaymentCode) (PaymentCode, error) {
					return createdPaymentCode, nil
				},
			}
			svc := NewService(repo)

			actual, err := svc.GeneratePaymentCode(reqBody)
			Expect(err).To(Not(HaveOccurred()))
			Expect(actual).To(Equal(createdPaymentCode))
		})

		It("Should failed for name not provided", func() {
			var code string = "test1234"
			reqBody := reqBodyPaymentCode{
				Name:        nil,
				PaymentCode: &code,
			}

			repo := fakeRepository{}
			svc := NewService(repo)

			_, err := svc.GeneratePaymentCode(reqBody)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(common.ErrIllegalArg))
		})

		It("Should failed for payment code not provided", func() {
			var name string = "TEST"
			reqBody := reqBodyPaymentCode{
				Name:        &name,
				PaymentCode: nil,
			}

			repo := fakeRepository{}
			svc := NewService(repo)

			_, err := svc.GeneratePaymentCode(reqBody)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(common.ErrIllegalArg))
		})
	})
})
