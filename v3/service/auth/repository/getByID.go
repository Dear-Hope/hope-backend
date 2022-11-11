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
		job, activities, photo, p.id AS profile_id
		FROM "auth".users AS u, "auth".profiles AS p
		WHERE u.id = p.user_id AND u.id=$1`,
		id,
	)
	if err != nil {
		log.Printf("user get by id: %s", err.Error())
		return nil, err
	}

	return &user, nil
}
