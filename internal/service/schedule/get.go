package schedule

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/schedule"
	"context"
	"net/http"
)

func (s *service) Get(ctx context.Context, expertId uint64) ([]schedule.Response, *response.ServiceError) {
	var (
		responses            []schedule.Response
		scheduleIds          []uint64
		mapScheduleTimeslots = map[uint64][]schedule.TimeslotResponse{}
	)

	schedules, err := s.repo.GetSchedulesByExpertId(ctx, expertId)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGetExpertScheduleFailed,
			Err:  err,
		}
	}

	for _, sch := range schedules {
		scheduleIds = append(scheduleIds, sch.Id)
	}
	timeslots, err := s.repo.GetTimeslotsByScheduleIds(ctx, scheduleIds)
	if err != nil {
		return nil, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorGetExpertScheduleFailed,
			Err:  err,
		}
	}

	for _, ts := range timeslots {
		tsResp := schedule.TimeslotResponse{
			Id:        ts.Id,
			StartTime: ts.StartTime,
			EndTime:   ts.EndTime,
			TypeIds:   ts.TypeIds,
		}
		if val, ok := mapScheduleTimeslots[ts.ScheduleId]; ok {
			val = append(val, tsResp)
			mapScheduleTimeslots[ts.ScheduleId] = val
		} else {
			mapScheduleTimeslots[ts.ScheduleId] = []schedule.TimeslotResponse{tsResp}
		}
	}

	for _, sch := range schedules {
		responses = append(responses, schedule.Response{
			Id:        sch.Id,
			Day:       sch.Day,
			IsActive:  sch.IsActive,
			Timeslots: mapScheduleTimeslots[sch.Id],
		})
	}

	return responses, nil
}
