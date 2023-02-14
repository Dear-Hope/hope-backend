package counseling

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/counseling/filter"
	"time"
)

type Repository interface {
	GetAllTopics() (model.Topics, error)
	StoreExpert(newExpert model.Expert, postIds []uint) (*model.Expert, error)
	GetAllExperts(f filter.ListExpert) (model.Experts, error)
	GetExpertById(id uint) (*model.Expert, error)
	UpdateExpert(updatedExpert model.Expert, updatedTopicIds, deletedTopicIds []uint) (*model.Expert, error)
	DeleteExpert(id uint) error
	GetExpertTopics(expertId uint) ([]uint, error)
	GetExpertUpcomingSchedule(expertId, typeId int64) (*model.Consultation, error)
	GetAllExpertSchedule(expertId, typeId int64, start, end time.Time) (model.Consultations, error)
	UpdateExpertScheduleUser(consulId, userId int64) (bool, error)
}
