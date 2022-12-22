package repository

import (
	"database/sql"
	"errors"
	"log"
)

func (ths *repository) GetTotalCommentsByPostID(postID uint) int {
	var totalComments int

	err := ths.db.Get(
		&totalComments,
		`SELECT count(id) FROM "storyroom".comments WHERE is_deleted = false AND post_id = $1`,
		postID,
	)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Printf("get post total comments: %s", err.Error())

		}
		return 0
	}

	return totalComments
}
