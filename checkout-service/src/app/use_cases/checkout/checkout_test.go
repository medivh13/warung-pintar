package checkout_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : warung-pintar/checkout-service
 */

import (
	"context"
	"errors"
	"testing"

	mockDTOitem "warung-pintar/checkout-service/mocks/app/dtos/checkout"
	mockService "warung-pintar/checkout-service/mocks/infra/service/item"

	dtoItem "warung-pintar/checkout-service/src/app/dtos/checkout"
	models "warung-pintar/checkout-service/src/infra/models"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUsecase struct {
	mock.Mock
}

type UsecaseCheckoutTest struct {
	suite.Suite
	checkoutService *mockService.MockItemsService

	usecase     CheckoutUsecaseInterface
	itemModels  []*models.Items
	dtoTest     *dtoItem.CheckoutReqDTO
	dtoTestFail *dtoItem.CheckoutReqDTO
	mockDTO     *mockDTOitem.MockCheckoutDTO
	itemResp    []*dtoItem.ItemDTO
	promoResp   []*dtoItem.PromoDTO
}

func (suite *UsecaseCheckoutTest) SetupTest() {
	suite.checkoutService = new(mockService.MockItemsService)
	suite.mockDTO = new(mockDTOitem.MockCheckoutDTO)
	suite.usecase = NewCheckoutUseCase(suite.checkoutService)

	suite.itemModels = []*models.Items{
		&models.Items{
			ID:    1,
			SKU:   "Test",
			Name:  "Test",
			Price: 100,
			Qty:   10,
		},
	}

	suite.dtoTest = &dtoItem.CheckoutReqDTO{
		Data: []*dtoItem.SkuReqDTO{
			&dtoItem.SkuReqDTO{
				Sku: "Test",
				Qty: 3,
			},
		},
	}

	suite.itemResp = []*dtoItem.ItemDTO{
		&dtoItem.ItemDTO{
			ID:    1,
			SKU:   "Test",
			Name:  "Test",
			Price: 100,
			Qty:   10,
		},
	}

	suite.promoResp = []*dtoItem.PromoDTO{
		&dtoItem.PromoDTO{
			ID:              1,
			SKU:             "Test",
			ValuePercentage: 0.1,
			MinimumQty:      10,
		},
	}

}

func (uc *UsecaseCheckoutTest) TestCheckoutSuccess() {
	var sku = []string{"Test"}
	uc.checkoutService.Mock.On("ListItems", sku).Return(uc.itemResp, nil)
	uc.checkoutService.Mock.On("ListPromos", sku).Return(uc.promoResp, nil)
	_, err := uc.usecase.Checkout(context.Background(), uc.dtoTest)
	uc.Equal(nil, err)
}

func (uc *UsecaseCheckoutTest) TestCheckoutFailRetrieveItems() {
	var sku = []string{"Test"}
	uc.checkoutService.Mock.On("ListItems", sku).Return(nil, errors.New(mock.Anything))

	_, err := uc.usecase.Checkout(context.Background(), uc.dtoTest)
	uc.Error(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UsecaseCheckoutTest))
}
