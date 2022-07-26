package mock_dto

import (
	dto "warung-pintar/promo-service/src/app/dtos/items"

	"github.com/stretchr/testify/mock"
)

type MockItemDTO struct {
	mock.Mock
}

func NewMockItemDTO() *MockItemDTO {
	return &MockItemDTO{}
}

var _ dto.ItemInterface = &MockItemDTO{}

func (m *MockItemDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
