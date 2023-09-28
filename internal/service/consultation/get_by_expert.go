package consultation

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/consultation"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/schedule"
	"context"
	"fmt"
	"net/http"
	"time"
)

func (s *service) GetByExpert(ctx context.Context, req consultation.ExpertListRequest) (
	*consultation.ExpertListResponse, *response.ServiceError) {
	var (
		listResponses    []consultation.ExpertResponse
		isFilterPerMonth = req.BookingMonth != ""
	)

	request := consultation.Consultation{
		ExpertId:    req.ExpertId,
		UserId:      req.UserId,
		BookingDate: req.BookingDate,
		Status:      req.Status.String(),
	}

	if isFilterPerMonth {
		request.BookingDate = req.BookingMonth
	}

	consultations, err := s.repo.GetAllConsultation(ctx, request, isFilterPerMonth)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGetConsultationFailed,
			Err:  err,
		}
	}

	for _, consul := range consultations {
		resp, err := s.constructExpertResponse(ctx, consul)
		if err != nil {
			return nil, err
		}
		listResponses = append(listResponses, *resp)
	}

	return &consultation.ExpertListResponse{
		Data:        listResponses,
		TotalClient: len(listResponses),
	}, nil
}

func (s *service) constructExpertResponse(ctx context.Context, consul consultation.Consultation) (
	*consultation.ExpertResponse, *response.ServiceError) {
	now := time.Now().UTC()

	usr, err := s.userRepo.GetUserById(ctx, consul.UserId)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGetUserFailed,
			Err:  err,
		}
	}
	date, err := time.Parse(time.RFC3339, consul.BookingDate+"T"+consul.StartTime+":00+07:00")
	timePeriod := "Sedang Berlangsung"
	if now.Before(date) || now.After(date) {
		timePeriod = fmt.Sprintf("%s, %d %s %d | %s - %s WIB", schedule.DayIndonesian[date.Weekday()], date.Day(),
			schedule.MonthsIndonesian[date.Month()-1], date.Year(), consul.StartTime, consul.EndTime)
	}
	return &consultation.ExpertResponse{
		Id:                  consul.Id,
		ClientName:          usr.Name,
		ClientPhoto:         usr.Photo,
		ClientNote:          consul.UserNotes,
		TypeId:              consul.TypeId,
		Status:              consultation.GetStatus(consul.Status).Text(),
		Time:                timePeriod,
		IsStartConsultation: now.After(date) && consul.Status == consultation.Accepted.String(),
	}, nil
}
