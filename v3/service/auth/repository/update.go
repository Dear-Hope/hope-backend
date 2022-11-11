package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) Update(user model.User) (*model.User, error) {
	rows, err := ths.db.NamedQuery(
		`WITH updated_query AS (UPDATE "auth".users
			SET email = :email, name = :name, alias = :alias
			WHERE id = :user_id RETURNING id
		)
		UPDATE "auth".profiles SET job = :job, activities = :activities,
		user_id = (SELECT id from updated_query)
		WHERE id = :profile_id RETURNING user_id, id AS profile_id`,
		user,
	)
	if err != nil {
		log.Printf("user update: %s", err.Error())
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
