package mock_cases

import (
	"context"

	dto "warung-pintar/checkout-service/src/app/dtos/checkout"

	"github.com/stretchr/testify/mock"
)

type MockCheckoutUseCase struct {
	mock.Mock
}

func (m *MockCheckoutUseCase) Checkout(ctx context.Context, checkoutDto *dto.CheckoutReqDTO) (float64, error) {
	args := m.Called(checkoutDto)
	var err error
	var total float64

	if n, ok := args.Get(0).(float64); ok {
		total = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return total, err
}
