package mock_repository

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : warung-pintar/checkout-service
 */

import (
	dto "warung-pintar/checkout-service/src/app/dtos/checkout"

	"github.com/stretchr/testify/mock"
)

type MockItemsService struct {
	mock.Mock
}

func (o *MockItemsService) ListItems(sku []string) ([]*dto.ItemDTO, error) {
	args := o.Called(sku)

	var (
		err      error
		respData []*dto.ItemDTO
	)
	if n, ok := args.Get(0).([]*dto.ItemDTO); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return respData, err
}

func (o *MockItemsService) ListPromos(sku []string) ([]*dto.PromoDTO, error) {
	args := o.Called(sku)

	var (
		err      error
		respData []*dto.PromoDTO
	)
	if n, ok := args.Get(0).([]*dto.PromoDTO); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return respData, err
}
