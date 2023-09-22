package schedule

import (
	"HOPE-backend/internal/entity/schedule"
	"context"
	"fmt"
)

func (r *repository) GetTimeslotsByDay(ctx context.Context, expertId, typeId uint64, day string) (schedule.Timeslots,
	error) {
	var timeslots schedule.Timeslots

	err := r.db.SelectContext(ctx, &timeslots,
		r.db.Rebind(`SELECT est.id, est.start_time, est.end_time 
						  FROM "expert".expert_schedule_timeslots est
						  LEFT JOIN "expert".expert_schedules es ON est.schedule_id = es.id
						  LEFT JOIN "expert".expert_timeslot_types ett ON est.id = ett.timeslot_id
                          WHERE es.expert_id = ? And es.day = ? AND ett.type_id = ? 
                          AND es.is_active = true AND est.is_deleted = false`),
		expertId, day, typeId)
	if err != nil {
		return nil, fmt.Errorf("[ScheduleRepo.GetTimeslotsByDay] Failed select: %w", err)
	}

	return timeslots, nil
}
