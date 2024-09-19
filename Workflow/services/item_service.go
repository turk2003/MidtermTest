package services

import (
	"github.com/turk2003/workflow/models"
	"github.com/turk2003/workflow/repositories"
)

type ItemService struct {
	Repository *repositories.ItemRepository
}

func NewItemService(repo *repositories.ItemRepository) *ItemService {
	return &ItemService{Repository: repo}
}

func (s *ItemService) CreateItem(item *models.Item) error {
	return s.Repository.Create(item)
}

func (s *ItemService) GetAllItems() ([]models.Item, error) {
	return s.Repository.FindAll()
}

func (s *ItemService) GetItemByID(id uint) (models.Item, error) {
	return s.Repository.FindByID(id)
}

func (s *ItemService) UpdateItem(item *models.Item) error {
	return s.Repository.Update(item)
}

func (s *ItemService) DeleteItem(id uint) error {
	return s.Repository.Delete(id)
}
