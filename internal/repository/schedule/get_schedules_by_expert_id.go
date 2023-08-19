package schedule

import (
	"HOPE-backend/internal/entity/schedule"
	"context"
	"fmt"
)

func (r *repository) GetSchedulesByExpertId(ctx context.Context, expertId uint64) (schedule.Schedules, error) {
	var schedules schedule.Schedules

	err := r.db.SelectContext(ctx, &schedules,
		r.db.Rebind(`SELECT id, day, is_active FROM "expert".
			expert_schedules WHERE expert_id = ? AND is_deleted = false`), expertId)
	if err != nil {
		return nil, fmt.Errorf("[ScheduleRepo.GetSchedulesByExpertId] Failed select: %w", err)
	}

	return schedules, nil
}
