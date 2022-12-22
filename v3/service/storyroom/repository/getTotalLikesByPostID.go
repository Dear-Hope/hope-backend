package repository

import (
	"database/sql"
	"errors"
	"log"
)

func (ths *repository) GetTotalLikesByPostID(postID uint) int {
	var totalLikes int

	err := ths.db.Get(
		&totalLikes,
		`SELECT count(id) FROM "storyroom".likes WHERE is_deleted = false AND post_id = $1`,
		postID,
	)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Printf("get post total likes: %s", err.Error())
		}
		return 0
	}

	return totalLikes
}
