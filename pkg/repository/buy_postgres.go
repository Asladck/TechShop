package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TechBuyPostgres struct {
	db *sqlx.DB
}

func NewTechBuyPostgres(db *sqlx.DB) *TechBuyPostgres {
	return &TechBuyPostgres{db: db}
}
func (r *TechBuyPostgres) GetPriceCart(userId string) (float64, error) {
	var total float64
	query := fmt.Sprintf(`SELECT COALESCE(SUM(i.price * ci.item_count),0) FROM %s ci INNER JOIN %s i ON ci.item_id = i.id WHERE ci.user_id = $1`, cartTable, itemsTable)
	err := r.db.Get(&total, query, userId)
	if err != nil {
		return 0, err
	}
	return total, err
}
func (r *TechBuyPostgres) BuyItem(userId, itemId string, stock int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var currentStock int
	checkStockQuery := fmt.Sprintf(`SELECT stock FROM %s WHERE id = $1 FOR UPDATE`, itemsTable)
	err = tx.QueryRow(checkStockQuery, itemId).Scan(&currentStock)
	if err != nil {
		return err
	}

	if currentStock < stock {
		return fmt.Errorf("недостаточно товара на складе")
	}
	insertOrderQuery := fmt.Sprintf(`
		INSERT INTO %s (user_id, item_id, item_count,status)
		VALUES ($1, $2, $3, $4)
	`, ordersTable)
	_, err = tx.Exec(insertOrderQuery, userId, itemId, stock, "pending")
	if err != nil {
		return err
	}
	updateStockQuery := fmt.Sprintf(`
		UPDATE %s SET stock = stock - $1 WHERE id = $2
	`, itemsTable)
	_, err = tx.Exec(updateStockQuery, stock, itemId)
	if err != nil {
		return err
	}

	return nil
}
