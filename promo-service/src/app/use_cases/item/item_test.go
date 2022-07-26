package item_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : warung-pintar/checkout-service
 */

import (
	"context"
	"errors"
	"testing"

	mockDTOitem "warung-pintar/promo-service/mocks/app/dtos/items"
	mockCheckoutRepo "warung-pintar/promo-service/mocks/domain/repository"

	dtoItem "warung-pintar/promo-service/src/app/dtos/items"
	models "warung-pintar/promo-service/src/infra/models"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUsecase struct {
	mock.Mock
}

type UsecaseCheckoutTest struct {
	suite.Suite
	checkoutrepo *mockCheckoutRepo.MockItemsRepo

	usecase     ItemUsecaseInterface
	itemModels  []*models.Items
	dtoTest     *dtoItem.ItemReqDTO
	dtoTestFail *dtoItem.ItemReqDTO
	mockDTO     *mockDTOitem.MockItemDTO
}

func (suite *UsecaseCheckoutTest) SetupTest() {
	suite.checkoutrepo = new(mockCheckoutRepo.MockItemsRepo)
	suite.mockDTO = new(mockDTOitem.MockItemDTO)
	suite.usecase = NewItemUseCase(suite.checkoutrepo)

	suite.itemModels = []*models.Items{
		&models.Items{
			ID:              1,
			SKU:             "Test",
			ValuePercentage: 0.1,
			MinimumQty:      10,
		},
	}

	suite.dtoTest = &dtoItem.ItemReqDTO{
		Data: []*dtoItem.DataItemReqDTO{
			&dtoItem.DataItemReqDTO{
				Sku: "Test",
			},
		},
	}

}

func (uc *UsecaseCheckoutTest) TestCheckoutSuccess() {
	var sku = []string{"Test"}
	uc.checkoutrepo.Mock.On("ListItems", sku).Return(uc.itemModels, nil)
	_, err := uc.usecase.ListItems(context.Background(), uc.dtoTest)
	uc.Equal(nil, err)
}

func (uc *UsecaseCheckoutTest) TestCheckoutFailRetrieveItems() {
	var sku = []string{"Test"}
	uc.checkoutrepo.Mock.On("ListItems", sku).Return(nil, errors.New(mock.Anything))

	_, err := uc.usecase.ListItems(context.Background(), uc.dtoTest)
	uc.Error(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UsecaseCheckoutTest))
}
