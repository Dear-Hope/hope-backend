package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := ths.db.Get(
		&user,
		`SELECT u.id AS user_id, email, password, name, alias, is_active,
		job, activities, photo, p.id AS profile_id, total_audio_played, total_time_played, longest_streak
		FROM `+user.TableWithSchemaName()+` AS u, `+user.Profile.TableWithSchemaName()+` AS p
		WHERE u.id = p.user_id AND u.id=$1`,
		id,
	)
	if err != nil {
		log.Printf("user get by id: %s", err.Error())
		return nil, err
	}

	return &user, nil
}
