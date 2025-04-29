package service

import (
	"TechShop/pkg/repository"
)

type TechBuyService struct {
	repo     repository.Buy
	repoItem repository.Item
}

func NewTechBuyService(repo repository.Buy, repoItem repository.Item) *TechBuyService {
	return &TechBuyService{repo: repo, repoItem: repoItem}
}

func (s *TechBuyService) GetPriceCart(userId string) (float64, error) {
	return s.repo.GetPriceCart(userId)
}
func (s *TechBuyService) BuyItem(userId, itemId string, stock int) error {
	if _, err := s.repoItem.GetItemById(itemId); err != nil {
		return err
	}
	return s.repo.BuyItem(userId, itemId, stock)
}
