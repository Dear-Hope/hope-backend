package consultation

import (
	"context"
	"fmt"
)

func (r *repository) GetMapBookedConsultation(ctx context.Context, expertId uint64, date string) (map[string]bool, error) {
	var (
		startTime, endTime string
		mapBookedConsul    = map[string]bool{}
	)
	rows, err := r.db.QueryxContext(ctx, r.db.Rebind(
		`SELECT start_time, end_time FROM "counsel".consultations 
			WHERE expert_id = ? AND booking_date = ? AND is_deleted = false`), expertId, date)
	if err != nil {
		return nil, fmt.Errorf("[ScheduleRepo.GetMapBookedConsultation] Failed select: %w", err)
	}

	for rows.Next() {
		err := rows.Scan(&startTime, &endTime)
		if err != nil {
			return nil, fmt.Errorf("[ScheduleRepo.GetMapBookedConsultation] Failed scan: %v", err)
		}
		mapBookedConsul[startTime+"_"+endTime] = true
	}

	return mapBookedConsul, nil
}
