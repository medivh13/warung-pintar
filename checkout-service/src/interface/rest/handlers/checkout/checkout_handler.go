package checkout_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : warung-pintar/checkout-service
 */

import (
	"encoding/json"
	"net/http"

	dtoCheckout "warung-pintar/checkout-service/src/app/dtos/checkout"
	usecases "warung-pintar/checkout-service/src/app/use_cases/checkout"
	common_error "warung-pintar/checkout-service/src/infra/errors"
	"warung-pintar/checkout-service/src/interface/rest/response"

	"gorm.io/gorm"
)

type CheckoutHandlerInterface interface {
	Checkout(w http.ResponseWriter, r *http.Request)
}

type checkoutHandler struct {
	response response.IResponseClient
	usecase  usecases.CheckoutUsecaseInterface
}

func NewCheckoutHandler(r response.IResponseClient, h usecases.CheckoutUsecaseInterface) CheckoutHandlerInterface {
	return &checkoutHandler{
		response: r,
		usecase:  h,
	}
}

func (h *checkoutHandler) Checkout(w http.ResponseWriter, r *http.Request) {

	postDTO := dtoCheckout.CheckoutReqDTO{}
	err := json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	err = postDTO.Validate()
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}

	data, err := h.usecase.Checkout(r.Context(), &postDTO)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			h.response.HttpError(w, common_error.NewError(common_error.STATUS_PAGE_NOT_FOUND, err))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_STATUS_PAGE, err))
		return
	}

	h.response.JSON(
		w,
		"Successful Checkout",
		data,
		nil,
	)
}
