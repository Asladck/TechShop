package service

import (
	"TechShop/models"
	"TechShop/pkg/repository"
)

type TechCartService struct {
	repo  repository.Cart
	irepo repository.Item
}

func NewTechCartService(repo repository.Cart, irepo repository.Item) *TechCartService {
	return &TechCartService{repo: repo, irepo: irepo}
}
func (s *TechCartService) AddToCart(userId, itemId string, countItem int) (string, error) {
	if _, err := s.irepo.GetItemById(itemId); err != nil {
		return "", err
	}
	return s.repo.AddToCart(userId, itemId, countItem)
}
func (s *TechCartService) GetCart(userId string) ([]models.Cart, error) {
	return s.repo.GetCart(userId)
}
func (s *TechCartService) GetCartItemById(userId, cartId string) (models.Cart, error) {
	return s.repo.GetCartItemById(userId, cartId)
}
func (s *TechCartService) Update(userId, cartId string, countItem int) error {
	if _, err := s.repo.GetCartItemById(userId, cartId); err != nil {
		return err
	}
	return s.repo.Update(userId, cartId, countItem)
}
func (s *TechCartService) Delete(userId, cartId string) error {
	if _, err := s.repo.GetCartItemById(userId, cartId); err != nil {
		return err
	}
	return s.repo.Delete(userId, cartId)
}
