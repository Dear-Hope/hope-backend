package schedule

import (
	"HOPE-backend/internal/entity/schedule"
	"HOPE-backend/pkg/helpers"
	"context"
	"fmt"
	"github.com/lib/pq"
	"strings"
	"time"
)

func (r *repository) UpdateTimeslots(ctx context.Context, req schedule.Timeslots) (*schedule.Timeslots, error) {
	var (
		insertTypeValue     string
		insertTimeslotValue string
		timeslotIds         []uint64
		now                 = time.Now().UTC()
	)

	for _, ts := range req {
		id := "default"
		if ts.Id != 0 {
			id = fmt.Sprint(ts.Id)
		}

		insertTimeslotValue += fmt.Sprintf("(%s, %d, '%s', '%s', '%s'::TIMESTAMP),",
			id, ts.ScheduleId, ts.StartTime, ts.EndTime, helpers.TimestampToStringFormat(now, time.RFC3339),
		)
	}

	query := fmt.Sprintf(`INSERT INTO "expert".expert_schedule_timeslots
    		(id, schedule_id, start_time, end_time, created_at)
			VALUES %s ON CONFLICT(id) DO UPDATE 
			SET start_time = excluded.start_time,
				end_time = excluded.end_time,
				updated_at = excluded.updated_at
			RETURNING id`,
		strings.TrimSuffix(insertTimeslotValue, ","),
	)

	rows, err := r.db.QueryxContext(ctx, r.db.Rebind(query))
	if err != nil {
		return nil, fmt.Errorf("[ScheduleRepo.UpdateTimeslots] Failed exec: %w", err)
	}

	i := 0
	for rows.Next() {
		if err := rows.Scan(&req[i].Id); err != nil {
			return nil, fmt.Errorf("[ScheduleRepo.UpdateTimeslots] Failed scan: %w", err)
		}
		i++
	}

	for _, ts := range req {
		for _, typeId := range ts.TypeIds {
			insertTypeValue += fmt.Sprintf("(%d, %d),", ts.Id, typeId)
		}

		timeslotIds = append(timeslotIds, ts.Id)
	}
	queryType := fmt.Sprintf(
		`WITH deleted AS (DELETE FROM "expert".expert_timeslot_types WHERE timeslot_id = ANY(?))
		INSERT INTO "expert".expert_timeslot_types(timeslot_id, type_id) VALUES%s`,
		strings.TrimSuffix(insertTypeValue, ","),
	)
	_, err = r.db.ExecContext(ctx, r.db.Rebind(queryType), pq.Array(timeslotIds))
	if err != nil {
		return nil, fmt.Errorf("[ScheduleRepo.UpdateTimeslots] Failed exec delete: %w", err)
	}

	return &req, nil
}
