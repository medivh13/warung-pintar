package usecases

import (
	itemUc "warung-pintar/item-service/src/app/use_cases/item"
)

type AllUseCases struct {
	ItemUseCase itemUc.ItemUsecaseInterface
}
