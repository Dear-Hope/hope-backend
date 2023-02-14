package repository

import (
	"log"
)

func (ths *repository) UpdateExpertScheduleUser(consulId, userId int64) (bool, error) {
	res, err := ths.db.Exec(
		ths.db.Rebind(
			`WITH consul AS (
				UPDATE "counseling".consultations SET user_id = ?
				FROM "counseling".consultations c JOIN "counseling".expert_schedules es
				ON c.schedule_id = es.id WHERE c.id = ? AND es.is_booked = false RETURNING c.schedule_id
			)
			UPDATE "counseling".expert_schedules SET is_booked = true WHERE id = (SELECT schedule_id FROM consul LIMIT 1)`,
		),
		userId, consulId,
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
