package service

import (
	"TechShop/models"
	"TechShop/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	GenerateToken(username, password, email string) (string, string, error)
	ParseRefToken(tokenR string) (string, error)
	ParseToken(token string) (string, error)
	GenerateAccToken(userId string) (string, error)
}
type Item interface {
	GetItems() ([]models.Item, error)
	GetItemById(itemId string) (models.Item, error)
	GetItemStock(itemId string) (int, error)
}
type Cart interface {
	AddToCart(userId, itemId string, countItem int) (string, error)
	GetCart(userId string) ([]models.Cart, error)
	GetCartItemById(userId, cartId string) (models.Cart, error)
	Update(userId, cartId string, countItem int) error
	Delete(userId, cartId string) error
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
	DeliveringOrder(userId, orderId string) error
	DeliveredOrder(userId, orderId string) error
}
type Buy interface {
	GetPriceCart(userId string) (float64, error)
	BuyItem(userId, itemId string, stock int) error
}

type Service struct {
	Authorization
	Cart
	Buy
	WishList
	Item
	Order
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
		Item:          NewTechItemService(rep.Item),
		WishList:      NewTechWishService(rep.WishList, rep.Item),
		Cart:          NewTechCartService(rep.Cart, rep.Item),
		Buy:           NewTechBuyService(rep.Buy, rep.Item),
		Order:         NewTechOrderService(rep.Order),
	}
}
