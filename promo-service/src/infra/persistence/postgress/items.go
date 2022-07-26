package postgres

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : warung-pintar/promo-service
 */

import (
	"context"

	repositories "warung-pintar/promo-service/src/domain/repositories"
	models "warung-pintar/promo-service/src/infra/models"

	"gorm.io/gorm"
)

type itemsRepository struct {
	connection *gorm.DB
}

func NewItemsRepository(db *gorm.DB) repositories.ItemsRepository {
	return &itemsRepository{
		connection: db,
	}
}

func (repo *itemsRepository) ListItems(ctx context.Context, sku []string) ([]*models.Items, error) {
	var itemsModel []*models.Items

	q := repo.connection.WithContext(ctx)
	if err := q.Raw(`SELECT id, sku, value_percentage, minimum_qty FROM warung_pintar.promos WHERE deleted_at IS NULL AND sku IN ?`, sku).Find(&itemsModel).Error; err != nil {
		return nil, err
	}

	return itemsModel, nil
}
