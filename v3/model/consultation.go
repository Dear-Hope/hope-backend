package model

import (
	"database/sql"
	"time"
)

type (
	Consultation struct {
		ID         int64         `db:"id"`
		ExpertId   int64         `db:"expert_id"`
		ScheduleId int64         `db:"schedule_id"`
		TypeId     int64         `db:"type_id"`
		UserId     sql.NullInt64 `db:"user_id"`
		StartAt    time.Time     `db:"start_at"`
		EndAt      time.Time     `db:"end_at"`
		IsBooked   bool          `db:"is_booked"`
	}

	Consultations []Consultation
)

func (ths Consultation) ToConsulResponse() *ConsultationResponse {
	return &ConsultationResponse{
		ID:         ths.ID,
		ScheduleId: ths.ScheduleId,
		StartAt:    ths.StartAt,
		EndAt:      ths.EndAt,
		IsBooked:   ths.IsBooked,
	}
}

func (ths Consultations) ToListConsulResponse() []ConsultationResponse {
	res := make([]ConsultationResponse, len(ths))
	for i, consul := range ths {
		res[i] = *consul.ToConsulResponse()
	}

	return res
}

type (
	ConsultationResponse struct {
		ID         int64     `json:"id"`
		ScheduleId int64     `json:"scheduleId"`
		StartAt    time.Time `json:"startAt"`
		EndAt      time.Time `json:"endAt"`
		IsBooked   bool      `json:"isBooked"`
	}

	BookingRequest struct {
		Id     int64 `json:"id"`
		UserId int64 `json:"userId"`
	}

	NewConsultationRequest struct {
		Schedules []ScheduleRequest `json:"schedules"`
	}

	ScheduleRequest struct {
		StartAt int64   `json:"startAt"`
		EndAt   int64   `json:"endAt"`
		Types   []int64 `json:"types"`
	}
)
