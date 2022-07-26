package repositories

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : warung-pintar/promo-service
 */

import (
	"context"

	models "warung-pintar/promo-service/src/infra/models"
)

type ItemsRepository interface {
	ListItems(context context.Context, sku []string) ([]*models.Items, error)
}
