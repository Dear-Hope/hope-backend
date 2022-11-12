package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := ths.db.Get(
		&user,
		`SELECT u.id AS user_id, email, password, name, alias, is_active, secret_key,
		job, activities, photo, p.id AS profile_id
		FROM `+user.TableWithSchemaName()+` AS u, `+model.Profile{}.TableWithSchemaName()+` AS p
		WHERE u.id = p.user_id AND email=$1`,
		email,
	)
	if err != nil {
		log.Printf("user get by email: %s", err.Error())
		return nil, err
	}

	return &user, nil
}
