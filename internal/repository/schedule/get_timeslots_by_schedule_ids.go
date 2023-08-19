package schedule

import (
	"HOPE-backend/internal/entity/schedule"
	"context"
	"fmt"
	"github.com/lib/pq"
)

func (r *repository) GetTimeslotsByScheduleIds(ctx context.Context, timeslotIds []uint64) (schedule.Timeslots, error) {
	var timeslots schedule.Timeslots

	err := r.db.SelectContext(ctx, &timeslots,
		r.db.Rebind(`SELECT est.id, est.schedule_id, est.start_time, est.end_time,
       		array_remove(array_agg(ett.type_id), NULL) as type_ids
			FROM "expert".expert_schedule_timeslots est
			LEFT JOIN expert.expert_schedules es on es.id = est.schedule_id
			LEFT JOIN "expert".expert_timeslot_types ett ON est.id = ett.timeslot_id
			WHERE es.id = ANY(?) GROUP BY est.id, est.start_time, est.end_time ORDER BY est.id`),
		pq.Array(timeslotIds),
	)
	if err != nil {
		return nil, fmt.Errorf("[ScheduleRepo.GetTimeslotsByScheduleIds] Failed select: %w", err)
	}

	return timeslots, nil
}
