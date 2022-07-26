package mock_cases

import (
	"context"

	dto "warung-pintar/promo-service/src/app/dtos/items"

	"github.com/stretchr/testify/mock"
)

type MockItemUseCase struct {
	mock.Mock
}

func (m *MockItemUseCase) ListItems(ctx context.Context, checkoutDto *dto.ItemReqDTO) ([]*dto.ItemRespDTO, error) {
	args := m.Called(checkoutDto)
	var (
		err      error
		respData []*dto.ItemRespDTO
	)
	if n, ok := args.Get(0).([]*dto.ItemRespDTO); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return respData, err
}
