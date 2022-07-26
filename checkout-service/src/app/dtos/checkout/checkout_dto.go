package items_dto

import validation "github.com/go-ozzo/ozzo-validation"

type CheckoutInterface interface {
	Validate() error
}

type CheckoutReqDTO struct {
	Data []*SkuReqDTO
}

type SkuReqDTO struct {
	Sku string `json:"sku"`
	Qty int64  `json:"qty"`
}

func (dto *CheckoutReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Data, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type ItemReqDTO struct {
	Data []*DataItemReqDTO `json:"data"`
}

type DataItemReqDTO struct {
	Sku string `json:"sku"`
}

type ItemsRespDTO struct {
	Data []*ItemDTO `json:"data"`
}

type ItemDTO struct {
	ID    int64   `json:"id"`
	SKU   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   float64 `json:"inventory_qty"`
}

type PromoRespDTO struct {
	Data []*PromoDTO `json:"data"`
}

type PromoDTO struct {
	ID              int64   `json:"id"`
	SKU             string  `json:"sku"`
	ValuePercentage float64 `json:"value_percentage"`
	MinimumQty      int64   `json:"minimum_qty"`
}
