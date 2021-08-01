package code

import (
	"testing"
	"time"

	"github.com/nadhiljanitra/nadhil-golang-training-beginner/common"
	"github.com/stretchr/testify/assert"
)

// Creating test using "testing" package

type arg struct {
	Name        *string `json:"name"`
	PaymentCode *string `json:"payment_code"`
}

var argName string = "test"
var argCodes string = "test1234"

var collections = []struct {
	in  arg
	out PaymentCode
}{
	{
		arg{&argName, &argCodes}, PaymentCode{
			Name:           "test",
			PaymentCode:    "test1234",
			Status:         "ACTIVE",
			CreatedAt:      time.Now().UTC().Truncate(1 * time.Second),
			UpdatedAt:      time.Now().UTC().Truncate(1 * time.Second),
			ExpirationDate: time.Now().UTC().AddDate(50, 0, 0).Format(time.RFC3339),
		},
	},
}

func TestValidData(t *testing.T) {
	for _, tt := range collections {
		t.Run("Succesfull build payment code", func(t *testing.T) {
			actual, err := NewPaymentCode(reqBodyPaymentCode(tt.in))
			assert.NoError(t, err)
			assert.Equal(t, tt.out, actual)
		})
	}
}

var failedData = []struct {
	in  arg
	out PaymentCode
}{
	{arg{&argName, nil}, PaymentCode{}},
	{arg{nil, nil}, PaymentCode{}},
}

func TestInValidData(t *testing.T) {
	for _, tt := range failedData {
		t.Run("failed build payment code", func(t *testing.T) {
			actual, err := NewPaymentCode(reqBodyPaymentCode(tt.in))
			assert.Equal(t, tt.out, actual)
			assert.Equal(t, err, common.ErrIllegalArg)
		})
	}
}
