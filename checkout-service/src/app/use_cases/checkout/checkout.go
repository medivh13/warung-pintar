package checkout_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : warung-pintar/checkout-service
 */

import (
	"context"
	"log"
	dtoCheckout "warung-pintar/checkout-service/src/app/dtos/checkout"
	msHelper "warung-pintar/checkout-service/src/infra/helpers"
	itemService "warung-pintar/checkout-service/src/infra/service/item"
)

type CheckoutUsecaseInterface interface {
	Checkout(ctx context.Context, checkoutDto *dtoCheckout.CheckoutReqDTO) (float64, error)
}

type chechkoutUseCase struct {
	ItemsService itemService.ItemServices
}

func NewCheckoutUseCase(itemsSrv itemService.ItemServices) *chechkoutUseCase {
	return &chechkoutUseCase{
		ItemsService: itemsSrv,
	}
}

func (uc *chechkoutUseCase) Checkout(ctx context.Context, checkoutDto *dtoCheckout.CheckoutReqDTO) (float64, error) {
	var sku []string

	for _, val := range checkoutDto.Data {
		sku = append(sku, val.Sku)
	}
	dataItems, err := uc.ItemsService.ListItems(sku)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	dataPromos, _ := uc.ItemsService.ListPromos(sku)

	total, err := msHelper.CheckPromo(checkoutDto, dataItems, dataPromos)

	return total, nil
}
