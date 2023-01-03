package repository

import "log"

func (ths *repository) CreateBlockedUser(userID, blockedUserID uint) error {
	_, err := ths.db.Queryx(
		`INSERT INTO "auth".blocked_users (user_id, blocked_user_id) VALUES ($1, $2)`,
		userID, blockedUserID,
	)
	if err != nil {
		log.Printf("create blocked user: %s", err.Error())
		return err
	}

	return nil
}
