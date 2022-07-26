package mock_dto

import (
	dto "warung-pintar/checkout-service/src/app/dtos/checkout"

	"github.com/stretchr/testify/mock"
)

type MockCheckoutDTO struct {
	mock.Mock
}

func NewMockCheckoutDTO() *MockCheckoutDTO {
	return &MockCheckoutDTO{}
}

var _ dto.CheckoutInterface = &MockCheckoutDTO{}

func (m *MockCheckoutDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
