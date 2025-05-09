package repository

import (
	"TechShop/models"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	GetUser(username, password, email string) (models.User, error)
}
type Cart interface {
	AddToCart(userId, itemId string, countItem int) (string, error)
	GetCart(userId string) ([]models.Cart, error)
	GetCartItemById(userId, cartId string) (models.Cart, error)
	Update(userId, cartId string, countItem int) error
	Delete(userId, cartId string) error
}
type Item interface {
	GetItems() ([]models.Item, error)
	GetItemById(itemId string) (models.Item, error)
	GetItemStock(itemId string) (int, error)
}
type WishList interface {
	AddToWishlist(userId, itemId string) (string, error)
	GetWishlist(userId string) ([]models.Item, error)
	DeleteWishItem(userId, itemId string) error
}
type Order interface {
	GetOrders(userId string) ([]models.Order, error)
	GetOrderById(userId, orderId string) (models.Order, error)
	CreateOrdersFromCart(userId string) error
	CreateOrderFromCart(userId, cartId string) error
	CancelOrder(userId, orderId string) error
	DeliveredOrder(userId, orderId string) error
	DeliveringOrder(userId, orderId string) error
}
type Buy interface {
	GetPriceCart(userId string) (float64, error)
	BuyItem(userId, itemId string, stock int) error
}
type Repository struct {
	Authorization
	Cart
	Item
	WishList
	Order
	Buy
}

func NewRepository(db *sqlx.DB, redis *redis.Client) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Item:          NewTechItemPostgres(db, redis),
		WishList:      NewWishlistPostgres(db, redis),
		Cart:          NewTechCartPostgres(db, redis),
		Buy:           NewTechBuyPostgres(db),
		Order:         NewTechOrderPostgres(db),
	}
}
