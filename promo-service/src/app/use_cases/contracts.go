package usecases

import (
	itemUc "warung-pintar/promo-service/src/app/use_cases/item"
)

type AllUseCases struct {
	ItemUseCase itemUc.ItemUsecaseInterface
}
