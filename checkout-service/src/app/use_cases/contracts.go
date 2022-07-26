package usecases

import (
	checkoutUc "warung-pintar/checkout-service/src/app/use_cases/checkout"
)

type AllUseCases struct {
	CheckoutUseCase checkoutUc.CheckoutUsecaseInterface
}
