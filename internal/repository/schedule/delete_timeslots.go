package schedule

import (
	"context"
	"fmt"
	"github.com/lib/pq"
)

func (r *repository) DeleteTimeslots(ctx context.Context, deletedIds []uint64) error {
	_, err := r.db.ExecContext(ctx,
		r.db.Rebind(`WITH deletedTypeIds AS (DELETE FROM "expert".expert_timeslot_types WHERE timeslot_id = ANY(?))
			DELETE FROM "expert".expert_schedule_timeslots WHERE id = ANY(?)`),
		pq.Array(deletedIds), pq.Array(deletedIds),
	)
	if err != nil {
		return fmt.Errorf("[ScheduleRepo.DeleteTimeslots] Failed exec: %w", err)
	}

	return nil
}
