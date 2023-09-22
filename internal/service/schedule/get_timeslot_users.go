package schedule

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/schedule"
	"context"
	"fmt"
	"net/http"
	"time"
)

func (s *service) GetTimeslotUsers(ctx context.Context, expertId, typeId uint64, date int64) (
	[]schedule.TimeslotUserResponse, *response.ServiceError) {

	var (
		responses []schedule.TimeslotUserResponse
		datetime  = time.UnixMilli(date)
	)

	timeslots, err := s.repo.GetTimeslotsByDay(ctx, expertId, typeId,
		schedule.DayIndonesian[datetime.Weekday()])
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGetExpertScheduleFailed,
			Err:  err,
		}
	}

	mapBookedConsul, err := s.consulRepo.GetMapBookedConsultation(ctx, expertId, datetime.Format(time.DateOnly))
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGetExpertScheduleFailed,
			Err:  err,
		}
	}

	for _, timeslot := range timeslots {
		start, _ := time.Parse("15:04", timeslot.StartTime)
		end, _ := time.Parse("15:04", timeslot.EndTime)
		for d := start; d != end; d = d.Add(time.Hour) {
			startTime := fmt.Sprintf("%d:00", d.Hour())
			endTime := fmt.Sprintf("%d:00", d.Hour()+1)
			responses = append(responses, schedule.TimeslotUserResponse{
				StartTime: startTime,
				EndTime:   endTime,
				IsBooked:  mapBookedConsul[startTime+"_"+endTime],
			})
		}
	}

	return responses, nil
}
