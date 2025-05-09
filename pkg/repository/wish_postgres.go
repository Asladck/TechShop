package repository

import (
	"TechShop/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"time"
)

type WishlistPostgres struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewWishlistPostgres(db *sqlx.DB, client *redis.Client) *WishlistPostgres {
	return &WishlistPostgres{db: db, redis: client}
}
func (r *WishlistPostgres) AddToWishlist(userId, itemId string) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (user_id, item_id) VALUES ($1,$2) RETURNING id", wishlistTable)
	err := r.db.QueryRow(query, userId, itemId).Scan(&id)
	if err != nil {
		return "", err
	}
	cacheKey := fmt.Sprintf("wishlist:user:%s", userId)
	_ = r.redis.Del(context.Background(), cacheKey).Err()
	return id, err
}

func (r *WishlistPostgres) GetWishlist(userId string) ([]models.Item, error) {
	var wishItems []models.Item
	cache := fmt.Sprintf("wishlist:user:%s", userId)
	cachedData, err := r.redis.Get(context.Background(), cache).Result()
	if err == nil {
		err = json.Unmarshal([]byte(cachedData), &wishItems)
		if err == nil {
			return wishItems, err
		}
	}
	query := fmt.Sprintf("SELECT i.id, i.title, i.description, i.price, i.stock, i.image_url, i.created_at AS item_created_at, i.updated_at AS item_updated_at FROM %s i INNER JOIN %s w ON i.id = w.item_id WHERE w.user_id = $1", itemsTable, wishlistTable)
	err = r.db.Select(&wishItems, query, userId)
	if err != nil {
		return nil, err
	}
	dataToCache, err := json.Marshal(wishItems)
	if err == nil {
		_ = r.redis.Set(context.Background(), cache, dataToCache, 10*time.Minute).Err()
	}
	return wishItems, err
}

func (r *WishlistPostgres) DeleteWishItem(userId, itemId string) error {
	query := fmt.Sprintf("DELETE FROM %s w WHERE user_id=$1 AND item_id=$2", wishlistTable)
	_, err := r.db.Exec(query, userId, itemId)
	if err != nil {
		return err
	}
	redisKey := fmt.Sprintf("wishlist:user:%s", userId)
	err = r.redis.Del(context.Background(), redisKey).Err()
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	return err
}
