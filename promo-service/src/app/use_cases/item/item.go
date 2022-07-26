package item_usecases

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : warung-pintar/promo-service
 */

import (
	"context"
	"log"
	dto "warung-pintar/promo-service/src/app/dtos/items"
	"warung-pintar/promo-service/src/domain/repositories"
)

type ItemUsecaseInterface interface {
	ListItems(ctx context.Context, checkoutDto *dto.ItemReqDTO) ([]*dto.ItemRespDTO, error)
}

type itemUseCase struct {
	ItemsRepo repositories.ItemsRepository
}

func NewItemUseCase(itemsRepo repositories.ItemsRepository) *itemUseCase {
	return &itemUseCase{
		ItemsRepo: itemsRepo,
	}
}

func (uc *itemUseCase) ListItems(ctx context.Context, checkoutDto *dto.ItemReqDTO) ([]*dto.ItemRespDTO, error) {
	var sku []string

	for _, val := range checkoutDto.Data {
		sku = append(sku, val.Sku)
	}
	dataItems, err := uc.ItemsRepo.ListItems(ctx, sku)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto.ToItems(dataItems), nil
}
