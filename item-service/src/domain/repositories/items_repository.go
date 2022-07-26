package repositories

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : warung-pintar/item-service
 */

import (
	"context"

	models "warung-pintar/item-service/src/infra/models"
)

type ItemsRepository interface {
	ListItems(context context.Context, sku []string) ([]*models.Items, error)
}
