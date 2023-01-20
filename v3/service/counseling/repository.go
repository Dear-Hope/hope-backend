package counseling

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/counseling/filter"
)

type Repository interface {
	GetAllTopics() (model.Topics, error)
	StoreExpert(newExpert model.Expert, postIds []uint) (*model.Expert, error)
	GetAllExperts(f filter.ListExpert) (model.Experts, error)
	GetExpertById(id uint) (*model.Expert, error)
	UpdateExpert(updatedExpert model.Expert, updatedTopicIds, deletedTopicIds []uint) (*model.Expert, error)
	DeleteExpert(id uint) error
	GetExpertTopics(expertId uint) ([]uint, error)
}
