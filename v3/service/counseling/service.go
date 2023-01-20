package counseling

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/counseling/filter"
)

type Service interface {
	ListTopics() ([]model.TopicResponse, *model.ServiceError)
	CreateExpert(req model.CreateUpdateExpertRequest) (*model.ExpertResponse, *model.ServiceError)
	ListExperts(f filter.ListExpert) ([]model.ExpertResponse, *model.ServiceError)
	GetExpert(id uint) (*model.ExpertResponse, *model.ServiceError)
	UpdateExpert(req model.CreateUpdateExpertRequest) (*model.ExpertResponse, *model.ServiceError)
	DeleteExpert(id uint) *model.ServiceError
}
