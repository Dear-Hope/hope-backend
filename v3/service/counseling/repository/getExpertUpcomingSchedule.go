package repository

import (
	"HOPE-backend/v3/model"
	"log"
	"time"
)

func (ths *repository) GetExpertUpcomingSchedule(expertId, typeId int64) (*model.Consultation, error) {
	var consul model.Consultation
	err := ths.db.Get(
		&consul,
		ths.db.Rebind(
			`SELECT c.id, c.schedule_id, es.start_at, es.end_at, es.is_booked FROM "counseling".expert_schedules es 
   					JOIN "counseling".consultations c ON es.id = c.schedule_id
    				WHERE es.expert_id = ? AND c.type_id = ? AND es.start_at > ? AND es.is_booked = false LIMIT 1`,
		),
		expertId, typeId, time.Now().UTC(),
	)
	if err != nil {
		log.Printf("get expert's upcoming schedule: %s", err.Error())
		return nil, err
	}

	return &consul, nil
}
