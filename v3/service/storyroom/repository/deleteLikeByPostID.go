package repository

import (
	"database/sql"
	"errors"
	"log"
)

func (ths *repository) DeleteLikeByPostID(postID uint) error {
	_, err := ths.db.Queryx(
		`UPDATE "storyroom".likes SET is_deleted = true, updated_at = now() WHERE post_id = $1`,
		postID,
	)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		log.Printf("delete like: %s", err.Error())
		return err
	}

	return nil
}
