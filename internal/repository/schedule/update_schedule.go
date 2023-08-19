package schedule

import (
	"HOPE-backend/internal/entity/schedule"
	"context"
	"fmt"
)

func (r *repository) UpdateSchedule(ctx context.Context, req schedule.Schedule) (*schedule.Schedule, error) {
	query := `UPDATE "expert".expert_schedules SET is_active = ?, updated_at = ? WHERE id = ? AND is_deleted = false`

	_, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.IsActive, req.UpdatedAt, req.Id)
	if err != nil {
		return nil, fmt.Errorf("[ScheduleRepo.UpdateSchedule] Failed exec: %w", err)
	}

	return &req, nil
}
