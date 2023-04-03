package repository

import (
	"HOPE-backend/v3/model"
	"fmt"
	"log"
	"strings"
	"time"
)

func (ths *repository) StoreExpertSchedule(schedules []model.ScheduleRequest, expertId int64) (int64, error) {
	var (
		insertScheduleValue     string
		insertConsultationValue string
	)
	for _, schedule := range schedules {
		startAt := time.UnixMilli(schedule.StartAt).UTC().Format(time.RFC3339)
		endAt := time.UnixMilli(schedule.EndAt).UTC().Format(time.RFC3339)
		insertScheduleValue += fmt.Sprintf(
			"(%d, '%v'::timestamp, '%v'::timestamp),", expertId, startAt, endAt,
		)
		for _, typeId := range schedule.Types {
			insertConsultationValue += fmt.Sprintf(
				"((SELECT id from schedules WHERE start_at = '%v'::timestamp), %d),", startAt, typeId,
			)
		}
	}

	insertScheduleQuery := fmt.Sprintf(`WITH schedules AS (
		INSERT INTO "counseling".expert_schedules(expert_id, start_at, end_at) VALUES 
		%s RETURNING id, start_at)`, strings.TrimSuffix(insertScheduleValue, ","),
	)

	insertConsultationQuery := fmt.Sprintf(`INSERT INTO "counseling".consultations(schedule_id, type_id) VALUES 
		%s;`, strings.TrimSuffix(insertConsultationValue, ","),
	)

	query := fmt.Sprintf(`%s %s`, insertScheduleQuery, insertConsultationQuery)

	result, err := ths.db.Exec(ths.db.Rebind(query))
	if err != nil {
		log.Printf("new expert schedule create failed: %s", err.Error())
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("new expert schedule create failed: %s", err.Error())
		return 0, err
	}

	return rowsAffected, nil
}
