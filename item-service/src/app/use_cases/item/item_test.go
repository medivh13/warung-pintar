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

	mockDTOitem "warung-pintar/item-service/mocks/app/dtos/items"
	mockCheckoutRepo "warung-pintar/item-service/mocks/domain/repository"

	dtoItem "warung-pintar/item-service/src/app/dtos/items"
	models "warung-pintar/item-service/src/infra/models"

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
			ID:    1,
			SKU:   "Test",
			Name:  "Test",
			Price: 100,
			Qty:   10,
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

func (uc *UsecaseCheckoutTest) TestListItemSuccess() {
	var sku = []string{"Test"}
	uc.checkoutrepo.Mock.On("ListItems", sku).Return(uc.itemModels, nil)
	_, err := uc.usecase.ListItems(context.Background(), uc.dtoTest)
	uc.Equal(nil, err)
}

func (uc *UsecaseCheckoutTest) TestListItemFailRetrieveItems() {
	var sku = []string{"Test"}
	uc.checkoutrepo.Mock.On("ListItems", sku).Return(nil, errors.New(mock.Anything))

	_, err := uc.usecase.ListItems(context.Background(), uc.dtoTest)
	uc.Error(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UsecaseCheckoutTest))
}
