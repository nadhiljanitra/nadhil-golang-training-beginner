package inquiry

import (
	"context"

	"github.com/google/uuid"
	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	Context("Get Inquiry by Transaction ID", func() {
		var createdInquiry Inquiry

		BeforeEach(func() {
			createdInquiry = Inquiry{
				ID:            uuid.New().String(),
				TransactionID: "transaction123",
				PaymentCode:   "test1234",
			}
		})

		It("Should find inquiry", func() {
			repo := fakeRepository{
				getInquiryByTransactionIDfn: func(ctx context.Context, transactionID string) (Inquiry, error) {
					return createdInquiry, nil
				},
			}
			svc := NewService(repo)

			actual, err := svc.GetInquiryByTransactionID("test1234")
			Expect(err).To(Not(HaveOccurred()))
			Expect(actual).To(Equal(createdInquiry))
		})

		It("Should failed for inquiry not found", func() {
			repo := fakeRepository{
				getInquiryByTransactionIDfn: func(ctx context.Context, transactionID string) (Inquiry, error) {
					return Inquiry{}, common.ErrNotFound
				},
			}

			svc := NewService(repo)
			_, err := svc.GetInquiryByTransactionID("not-exist")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(common.ErrNotFound))
		})
	})
})
