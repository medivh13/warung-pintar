package checkout_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : warung-pintar/item-service
 */

import (
	"encoding/json"
	"log"
	"net/http"

	dtoCheckout "warung-pintar/item-service/src/app/dtos/items"
	usecases "warung-pintar/item-service/src/app/use_cases/item"
	common_error "warung-pintar/item-service/src/infra/errors"
	"warung-pintar/item-service/src/interface/rest/response"

	"gorm.io/gorm"
)

type ItemHandlerInterface interface {
	GetItems(w http.ResponseWriter, r *http.Request)
}

type itemHandler struct {
	response response.IResponseClient
	usecase  usecases.ItemUsecaseInterface
}

func NewItemHandler(r response.IResponseClient, h usecases.ItemUsecaseInterface) ItemHandlerInterface {
	return &itemHandler{
		response: r,
		usecase:  h,
	}
}

func (h *itemHandler) GetItems(w http.ResponseWriter, r *http.Request) {

	postDTO := dtoCheckout.ItemReqDTO{}
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

	data, err := h.usecase.ListItems(r.Context(), &postDTO)
	log.Println(data, "here data")
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
		"Successful Get Data",
		data,
		nil,
	)
}
