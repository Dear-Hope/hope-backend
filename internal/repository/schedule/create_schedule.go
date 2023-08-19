package schedule

import (
	"HOPE-backend/internal/entity/schedule"
	"context"
	"fmt"
	"strings"
)

func (r *repository) CreateSchedule(ctx context.Context, expertId uint64) error {
	var insertValue string
	for _, day := range schedule.Days {
		insertValue += fmt.Sprintf("(%d, '%s'),", expertId, day)
	}

	_, err := r.db.ExecContext(ctx,
		fmt.Sprintf(
			`INSERT INTO "expert".expert_schedules (expert_id, day) VALUES %s`,
			strings.TrimSuffix(insertValue, ",")),
	)
	if err != nil {
		return fmt.Errorf("[ScheduleRepo.CreateSchedule] Failed exec: %w", err)
	}

	return nil
}
