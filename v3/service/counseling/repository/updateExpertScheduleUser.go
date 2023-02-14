package repository

import (
	"log"
)

func (ths *repository) UpdateExpertScheduleUser(consulId, userId int64) (bool, error) {
	res, err := ths.db.Exec(
		ths.db.Rebind(
			`WITH row AS (
    			SELECT es.is_booked
    			FROM "counseling".expert_schedules es, "counseling".consultations c
    			WHERE c.schedule_id = es.id AND c.id = ?
			), consul AS (
    			UPDATE "counseling".consultations SET user_id = ?
    			WHERE id = ? AND (SELECT is_booked FROM row) = false
    			RETURNING schedule_id
			)
			UPDATE "counseling".expert_schedules SET is_booked = true WHERE id = (SELECT schedule_id FROM consul LIMIT 1);`,
		),
		consulId, userId, consulId,
	)
	if err != nil {
		log.Printf("expert schedule update user failed: %s", err.Error())
		return false, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return affected > 0, nil
}
