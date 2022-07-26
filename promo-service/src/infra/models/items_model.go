package models

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : warung-pintar/checkout-service
 */

import "time"

type Items struct {
	ID              int64     `gorm:"id"`
	SKU             string    `gorm:"sku"`
	ValuePercentage float64   `gorm:"value_percentage"`
	MinimumQty      float64   `gorm:"minimum_qty"`
	CreatedAt       time.Time `gorm:"created_at"`
	UpdatedAt       time.Time `gorm:"updated_at"`
}
