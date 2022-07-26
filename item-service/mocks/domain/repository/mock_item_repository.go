package mock_repository

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : warung-pintar/checkout-service
 */

import (
	"context"

	models "warung-pintar/item-service/src/infra/models"

	"github.com/stretchr/testify/mock"
)

type MockItemsRepo struct {
	mock.Mock
}

func (o *MockItemsRepo) ListItems(context context.Context, sku []string) ([]*models.Items, error) {
	args := o.Called(sku)

	var (
		err      error
		respData []*models.Items
	)
	if n, ok := args.Get(0).([]*models.Items); ok {
		respData = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return respData, err
}
