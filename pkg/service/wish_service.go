package service

import (
	"TechShop/models"
	"TechShop/pkg/repository"
)

type TechWishService struct {
	rep  repository.WishList
	irep repository.Item
}

func NewTechWishService(rep repository.WishList, irep repository.Item) *TechWishService {
	return &TechWishService{rep: rep, irep: irep}
}
func (r *TechWishService) AddToWishlist(userId, itemId string) (string, error) {
	if _, err := r.irep.GetItemById(itemId); err != nil {
		return "", err
	}
	return r.rep.AddToWishlist(userId, itemId)
}
func (r *TechWishService) GetWishlist(userId string) ([]models.Item, error) {
	return r.rep.GetWishlist(userId)
}
func (r *TechWishService) DeleteWishItem(userId, itemId string) error {
	return r.rep.DeleteWishItem(userId, itemId)
}
