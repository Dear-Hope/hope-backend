package repository

import (
	"database/sql"
	"errors"
	"log"
)

func (ths *repository) CheckIsPostLikedByUserID(postID, userID uint) bool {
	var likeID uint
	err := ths.db.Get(
		&likeID,
		`SELECT id FROM "storyroom".likes WHERE is_deleted = false AND post_id = $1 AND user_id = $2`,
		postID, userID,
	)
	if err != nil || likeID <= 0 {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Printf("get post total likes: %v", err)
		}
		return false
	}

	return true
}
