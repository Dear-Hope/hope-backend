package selfcare

import (
	"HOPE-backend/models"
)

type service struct {
	repo models.SelfCareRepository
}

func NewSelfCareService(repo models.SelfCareRepository) models.SelfCareService {
	return &service{
		repo: repo,
	}
}

func (ths *service) NewItem(req models.NewSelfCareItemRequest) (*models.SelfCareItem, error) {
	newItem := models.SelfCareItem{
		Mood:        req.Mood,
		Link:        req.Link,
		Description: req.Description,
		Title:       req.Title,
		Type:        req.Type,
	}

	item, err := ths.repo.Create(newItem)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (ths *service) GetItemsByMood(mood string) ([]*models.SelfCareItem, error) {
	items, err := ths.repo.GetItemsByMood(mood)
	if err != nil {
		return nil, err
	}

	return items, err
}
