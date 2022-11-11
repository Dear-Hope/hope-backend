package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths repository) Create(user model.User) (*model.User, error) {
	rows, err := ths.db.NamedQuery(
		`WITH new_user AS (INSERT INTO "auth".users (email, password, first_name, last_name, is_active, secret_key) 
		VALUES (:email, :password, :first_name, :last_name, :is_active, :secret_key) RETURNING id)
		INSERT INTO "auth".profiles (job, activities, photo, user_id)
		VALUES (:job, :activities, :photo, (SELECT id from new_user))
		RETURNING user_id, id as profile_id`,
		user,
	)
	if err != nil {
		log.Printf("repository - user create failed: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.ProfileID)
		if err != nil {
			log.Printf("repository - user create failed: %s", err.Error())
			return nil, err
		}
	}
	return &user, nil
}
