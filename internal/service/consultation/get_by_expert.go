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
	now := time.Now()

	usr, err := s.userRepo.GetUserById(ctx, consul.UserId)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGetUserFailed,
			Err:  err,
		}
	}
	start, err := time.Parse(time.RFC3339, consul.BookingDate+"T"+consul.StartTime+":00+07:00")
	end, err := time.Parse(time.RFC3339, consul.BookingDate+"T"+consul.EndTime+":00+07:00")
	dt := ""
	t := "Sedang Berlangsung"
	if now.Before(start.UTC()) || now.After(end.UTC()) {
		dt = fmt.Sprintf("%s, %d %s %d", schedule.DayIndonesian[start.Weekday()], start.Day(),
			schedule.MonthsIndonesian[start.Month()-1], start.Year())
		t = fmt.Sprintf("%s - %s WIB", consul.StartTime, consul.EndTime)
	}
	return &consultation.ExpertResponse{
		Id:                  consul.Id,
		ClientName:          usr.Name,
		ClientPhoto:         usr.Photo,
		ClientNote:          consul.UserNotes,
		TypeId:              consul.TypeId,
		Status:              consultation.GetStatus(consul.Status).Text(),
		Date:                dt,
		Time:                t,
		IsStartConsultation: now.After(start) && now.Before(end) && consul.Status == consultation.Accepted.String(),
		CounselNote:         consul.CounselNotes,
		Document:            consul.Document,
	}, nil
}
