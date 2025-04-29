package service

import (
	"TechShop/models"
	"TechShop/pkg/repository"
)

type TechItemService struct {
	repo repository.Item
}

func NewTechItemService(repo repository.Item) *TechItemService {
	return &TechItemService{repo: repo}
}
func (s *TechItemService) GetItems() ([]models.Item, error) {
	return s.repo.GetItems()
}
func (s *TechItemService) GetItemById(itemId string) (models.Item, error) {
	return s.repo.GetItemById(itemId)
}
func (s *TechItemService) GetItemStock(itemId string) (int, error) {
	return s.repo.GetItemStock(itemId)
}
