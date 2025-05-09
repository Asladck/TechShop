package repository

import (
	"TechShop/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"time"
)

type TechItemPostgres struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewTechItemPostgres(db *sqlx.DB, client *redis.Client) *TechItemPostgres {
	return &TechItemPostgres{db: db, redis: client}
}

func (r *TechItemPostgres) GetItems() ([]models.Item, error) {
	var items []models.Item
	cache := "items:all"
	cachedData, err := r.redis.Get(context.Background(), cache).Result()
	if err == nil {
		err = json.Unmarshal([]byte(cachedData), &items)
		if err == nil {
			return items, nil
		}
	}
	query := fmt.Sprintf("SELECT it.id, it.title, it.description, it.price, it.stock, it.image_url,it.created_at AS item_created_at, it.updated_at AS item_updated_at FROM %s it", itemsTable)
	err = r.db.Select(&items, query)
	if err != nil {
		return nil, err
	}

	dataToCache, err := json.Marshal(items)
	if err == nil {
		_ = r.redis.Set(context.Background(), cache, dataToCache, 10*time.Minute).Err()
	}
	return items, err
}
func (r *TechItemPostgres) GetItemById(itemId string) (models.Item, error) {
	var item models.Item

	cache := fmt.Sprintf("item:%s", itemId)
	cachedData, err := r.redis.Get(context.Background(), cache).Result()
	if err == nil {
		err = json.Unmarshal([]byte(cachedData), &item)
		if err == nil {
			return item, nil
		}
	}
	query := fmt.Sprintf("SELECT id, title, description, price, stock, image_url, created_at AS item_created_at, updated_at AS item_updated_at FROM %s WHERE id=$1", itemsTable)
	err = r.db.Get(&item, query, itemId)
	if err != nil {
		return item, err
	}

	dataToCache, err := json.Marshal(item)
	if err == nil {
		_ = r.redis.Set(context.Background(), cache, dataToCache, 10*time.Minute).Err()
	}
	return item, nil
}
func (r *TechItemPostgres) GetItemStock(itemId string) (int, error) {

	var stock int

	cache := fmt.Sprintf("item:%s", itemId)
	cachedData, err := r.redis.Get(context.Background(), cache).Result()
	if err == nil {
		err = json.Unmarshal([]byte(cachedData), &stock)
		if err == nil {
			return stock, err
		}
	}

	query := fmt.Sprintf(`SELECT item_count FROM %s WHERE item_id=$1`, itemsTable)
	err = r.db.Get(&stock, query, itemId)
	if err != nil {
		return 0, err
	}
	dataToCache, err := json.Marshal(stock)
	if err == nil {
		_ = r.redis.Set(context.Background(), cache, dataToCache, 10*time.Minute).Err()
	}
	return stock, err
}
