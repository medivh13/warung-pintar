package items_dto

import (
	models "warung-pintar/promo-service/src/infra/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ItemInterface interface {
	Validate() error
}

type ItemReqDTO struct {
	Data []*DataItemReqDTO
}

type DataItemReqDTO struct {
	Sku string `json:"sku"`
}

func (dto *ItemReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Data, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type ItemRespDTO struct {
	ID              int64   `json:"id"`
	SKU             string  `json:"sku"`
	ValuePercentage float64 `json:"value_percentage"`
	MinimumQty      float64 `json:"minimum_qty"`
}

func ToGetItem(d *models.Items) *ItemRespDTO {
	return &ItemRespDTO{
		ID:              d.ID,
		SKU:             d.SKU,
		ValuePercentage: d.ValuePercentage,
		MinimumQty:      d.MinimumQty,
	}
}

func ToItems(d []*models.Items) []*ItemRespDTO {
	var data []*ItemRespDTO
	for _, val := range d {
		data = append(data, ToGetItem(val))
	}
	return data
}
