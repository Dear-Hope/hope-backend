package schedule

import (
	"HOPE-backend/internal/constant"
	"HOPE-backend/internal/entity/response"
	"HOPE-backend/internal/entity/schedule"
	"context"
	"net/http"
	"time"
)

func (s *service) Update(ctx context.Context, req schedule.UpdateRequest) (bool, *response.ServiceError) {
	now := time.Now().UTC()

	// Get old schedules
	oldSchedules, svcErr := s.Get(ctx, req.ExpertId)
	if svcErr != nil {
		return false, &response.ServiceError{
			Code: http.StatusInternalServerError,
			Msg:  constant.ErrorUpdateExpertScheduleFailed,
			Err:  svcErr.Err,
		}
	}

	// Update schedule & timeslots
	for i, sch := range req.Schedules {
		// Update schedule only when new value != old value
		// i.e. when it's "truly" updated
		if sch.IsActive != oldSchedules[i].IsActive {
			_, err := s.repo.UpdateSchedule(ctx, schedule.Schedule{
				Id:        sch.Id,
				IsActive:  sch.IsActive,
				UpdatedAt: now,
			})
			if err != nil {
				return false, &response.ServiceError{
					Code: http.StatusInternalServerError,
					Msg:  constant.ErrorUpdateExpertScheduleFailed,
					Err:  err,
				}
			}
		}

		// Only update the timeslots when schedule is set to active
		// We don't care about the timeslots when the schedule is not active
		if sch.IsActive {
			// Check which timeslot that should be deleted
			// by checking the discrepancy between new timeslots and old timeslots
			deletedTimeslots := getDeletedTimeslots(sch.Timeslots, oldSchedules[i].Timeslots)
			if len(deletedTimeslots) > 0 {
				err := s.repo.DeleteTimeslots(ctx, deletedTimeslots)
				if err != nil {
					return false, &response.ServiceError{
						Code: http.StatusInternalServerError,
						Msg:  constant.ErrorUpdateExpertScheduleFailed,
						Err:  err,
					}
				}
			}

			// Update the remaining timeslots
			_, err := s.repo.UpdateTimeslots(ctx, sch.Timeslots.ToTimeslots(sch.Id))
			if err != nil {
				return false, &response.ServiceError{
					Code: http.StatusInternalServerError,
					Msg:  constant.ErrorUpdateExpertScheduleFailed,
					Err:  err,
				}
			}
		}
	}

	return true, nil
}

func getDeletedTimeslots(newTs []schedule.UpdateTimeslotRequest, oldTs []schedule.TimeslotResponse) []uint64 {
	deletedTimeslot := make([]uint64, 0)
	mapCheck := map[uint64]struct{}{}
	for _, nts := range newTs {
		mapCheck[nts.Id] = struct{}{}
	}
	for _, ots := range oldTs {
		if _, found := mapCheck[ots.Id]; !found {
			deletedTimeslot = append(deletedTimeslot, ots.Id)
		}
	}

	return deletedTimeslot
}
