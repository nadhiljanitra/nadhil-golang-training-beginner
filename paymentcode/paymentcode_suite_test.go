package code_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//Bootsraping a Suite for ginkgo https://onsi.github.io/ginkgo/#bootstrapping-a-suite
func TestPaymentcode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Paymentcode Suite")
}
