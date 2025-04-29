package repository

import (
	"TechShop/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CartPostgres struct {
	db *sqlx.DB
}

func NewTechCartPostgres(db *sqlx.DB) *CartPostgres {
	return &CartPostgres{db: db}
}
func (r *CartPostgres) AddToCart(userId, itemId string, countItem int) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (user_id, item_id, item_count) VALUES($1,$2,$3) RETURNING id", cartTable)
	err := r.db.QueryRow(query, userId, itemId, countItem).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, err
}
func (r *CartPostgres) GetCart(userId string) ([]models.Cart, error) {
	var cart []models.Cart
	query := fmt.Sprintf(`
		SELECT 
			ci.id AS cart_id,
			ci.user_id,
			ci.item_id AS cart_item_id,
			ci.item_count,
			ci.created_at AS cart_created_at,
			ci.updated_at AS cart_updated_at,
			i.id,
			i.title,
			i.description,
			i.price,
			i.stock,
			i.image_url,
			i.created_at AS item_created_at,
			i.updated_at AS item_updated_at
		FROM %s ci
		INNER JOIN %s i ON ci.item_id = i.id
		WHERE ci.user_id = $1
	`, cartTable, itemsTable)

	err := r.db.Select(&cart, query, userId)
	if err != nil {
		return nil, err
	}
	return cart, err
}
func (r *CartPostgres) GetCartItemById(userId, cartId string) (models.Cart, error) {
	var cart models.Cart
	query := fmt.Sprintf(`
	SELECT
		ci.id AS cart_id,
		ci.user_id,
		ci.item_id AS cart_item_id,
		ci.item_count,
		ci.created_at AS cart_created_at,
		ci.updated_at AS cart_updated_at,
		i.id ,
		i.title,
		i.description,
		i.price,
		i.stock,
		i.image_url,
		i.created_at AS item_created_at,
		i.updated_at AS item_updated_at
	FROM %s ci
	INNER JOIN %s i ON ci.item_id = i.id
	WHERE ci.user_id = $1 AND ci.id = $2
	`, cartTable, itemsTable)
	err := r.db.Get(&cart, query, userId, cartId)
	if err != nil {
		return cart, err
	}
	return cart, err
}
func (r *CartPostgres) Update(userId, cartId string, countItem int) error {
	query := fmt.Sprintf(`UPDATE %s ct SET item_count=$1 WHERE ct.user_id=$2 AND ct.id=$3`, cartTable)
	_, err := r.db.Exec(query, countItem, userId, cartId)
	if err != nil {
		return err
	}
	return nil
}
func (r *CartPostgres) Delete(userId, cartId string) error {
	query := fmt.Sprintf(`DELETE FROM %s ct WHERE ct.id=$1 AND ct.user_id=$2`, cartTable)
	_, err := r.db.Exec(query, cartId, userId)
	if err != nil {
		return err
	}
	return nil
}
