package repository

import (
	"HOPE-backend/v3/model"
	"log"
	"time"
)

func (ths *repository) GetAllExpertSchedule(expertId, typeId int64, start, end time.Time) (model.Consultations, error) {
	var consuls model.Consultations
	err := ths.db.Select(
		&consuls,
		ths.db.Rebind(
			`SELECT c.id, c.schedule_id, es.start_at, es.end_at, es.is_booked FROM "counseling".expert_schedules es 
   					JOIN "counseling".consultations c ON es.id = c.schedule_id
    				WHERE es.expert_id = ? AND c.type_id = ? AND es.start_at >= ? AND es.end_at <= ?`,
		),
		expertId, typeId, start, end,
	)
	if err != nil {
		log.Printf("get expert's schedule list: %s", err.Error())
		return nil, err
	}

	return consuls, nil
}
