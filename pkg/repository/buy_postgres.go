package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"time"
)

type TechBuyPostgres struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewTechBuyPostgres(db *sqlx.DB, client *redis.Client) *TechBuyPostgres {
	return &TechBuyPostgres{db: db, redis: client}
}
func (r *TechBuyPostgres) GetPriceCart(userId string) (float64, error) {
	var total float64
	cacheKey := fmt.Sprintf("cart_price:%s", userId)
	cachedData, err := r.redis.Get(context.Background(), cacheKey).Result()
	if err == nil {
		err = json.Unmarshal([]byte(cachedData), &total)
		if err == nil {
			return total, err
		}
	}
	query := fmt.Sprintf(`SELECT COALESCE(SUM(i.price * ci.item_count),0) FROM %s ci INNER JOIN %s i ON ci.item_id = i.id WHERE ci.user_id = $1`, cartTable, itemsTable)
	err = r.db.Get(&total, query, userId)
	if err != nil {
		return 0, err
	}
	dataToCache, err := json.Marshal(total)
	if err == nil {
		_ = r.redis.Set(context.Background(), cacheKey, dataToCache, 10*time.Minute).Err()
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
	cacheKey := fmt.Sprintf("cart_price:%s", userId)
	err = r.redis.Del(context.Background(), cacheKey).Err()
	if err != nil {
		fmt.Println("Ошибка при удалении кэшированного значения корзины:", err)
	}
	return nil
}
