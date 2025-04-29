package service

import (
	"TechShop/models"
	"TechShop/pkg/repository"
)

type TechOrderService struct {
	repo repository.Order
}

func NewTechOrderService(repo repository.Order) *TechOrderService {
	return &TechOrderService{repo: repo}
}
func (s *TechOrderService) GetOrders(userId string) ([]models.Order, error) {
	return s.repo.GetOrders(userId)
}
func (s *TechOrderService) GetOrderById(userId, orderId string) (models.Order, error) {
	return s.repo.GetOrderById(userId, orderId)
}
func (s *TechOrderService) CreateOrdersFromCart(userId string) error {
	return s.repo.CreateOrdersFromCart(userId)
}
func (s *TechOrderService) CreateOrderFromCart(userId, cartId string) error {
	return s.repo.CreateOrderFromCart(userId, cartId)
}
func (s *TechOrderService) CancelOrder(userId, orderId string) error {
	if _, err := s.repo.GetOrderById(userId, orderId); err != nil {
		return err
	}
	return s.repo.CancelOrder(userId, orderId)
}
func (s *TechOrderService) DeliveringOrder(userId, orderId string) error {
	if _, err := s.repo.GetOrderById(userId, orderId); err != nil {
		return err
	}
	return s.repo.DeliveringOrder(userId, orderId)
}
func (s *TechOrderService) DeliveredOrder(userId, orderId string) error {
	if _, err := s.repo.GetOrderById(userId, orderId); err != nil {
		return err
	}
	return s.repo.DeliveredOrder(userId, orderId)
}
