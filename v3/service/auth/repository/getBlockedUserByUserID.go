package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetBlockedUserByUserID(userID uint) (model.BlockAccounts, error) {
	var blockedUsers model.BlockAccounts
	err := ths.db.Select(
		&blockedUsers,
		`SELECT id, user_id, blocked_user_id FROM "auth".blocked_users WHERE is_deleted = false AND user_id = $1`,
		userID,
	)
	if err != nil {
		log.Printf("user get blocked user: %s", err.Error())
		return nil, err
	}

	return blockedUsers, nil
}
