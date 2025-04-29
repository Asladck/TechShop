package repository

import (
	"TechShop/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type WishlistPostgres struct {
	db *sqlx.DB
}

func NewWishlistPostgres(db *sqlx.DB) *WishlistPostgres {
	return &WishlistPostgres{db: db}
}
func (r *WishlistPostgres) AddToWishlist(userId, itemId string) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (user_id, item_id) VALUES ($1,$2) RETURNING id", wishlistTable)
	err := r.db.QueryRow(query, userId, itemId).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, err
}

func (r *WishlistPostgres) GetWishlist(userId string) ([]models.Item, error) {
	var wishItems []models.Item
	query := fmt.Sprintf("SELECT i.id, i.title, i.description, i.price, i.stock, i.image_url, i.created_at, i.updated_at FROM %s i INNER JOIN %s w ON i.id = w.item_id WHERE w.user_id = $1", itemsTable, wishlistTable)
	err := r.db.Select(&wishItems, query, userId)
	if err != nil {
		return nil, err
	}
	return wishItems, err
}

func (r *WishlistPostgres) DeleteWishItem(userId, itemId string) error {
	query := fmt.Sprintf("DELETE FROM %s w WHERE user_id=$1 AND item_id=$2", wishlistTable)
	_, err := r.db.Exec(query, userId, itemId)
	if err != nil {
		return err
	}
	return err
}
