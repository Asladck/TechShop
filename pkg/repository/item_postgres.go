package repository

import (
	"TechShop/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TechItemPostgres struct {
	db *sqlx.DB
}

func NewTechItemPostgres(db *sqlx.DB) *TechItemPostgres {
	return &TechItemPostgres{db: db}
}

func (r *TechItemPostgres) GetItems() ([]models.Item, error) {
	var items []models.Item
	query := fmt.Sprintf("SELECT it.id, it.title, it.description, it.price, it.stock, it.image_url,it.created_at AS item_created_at, it.updated_at AS item_updated_at FROM %s it", itemsTable)
	err := r.db.Select(&items, query)
	if err != nil {
		return nil, err
	}
	return items, err
}
func (r *TechItemPostgres) GetItemById(itemId string) (models.Item, error) {
	var item models.Item
	query := fmt.Sprintf("SELECT id, title, description, price, stock, image_url, created_at AS item_created_at, updated_at AS item_updated_at FROM %s WHERE id=$1", itemsTable)
	err := r.db.Get(&item, query, itemId)
	if err != nil {
		return item, err
	}
	return item, nil
}
func (r *TechItemPostgres) GetItemStock(itemId string) (int, error) {
	var stock int
	query := fmt.Sprintf(`SELECT item_count FROM %s WHERE item_id=$1`, itemsTable)
	err := r.db.Get(&stock, query, itemId)
	if err != nil {
		return 0, err
	}
	return stock, err
}
