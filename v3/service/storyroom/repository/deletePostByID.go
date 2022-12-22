package repository

import (
	"database/sql"
	"errors"
	"log"
)

func (ths *repository) DeletePostByID(postID uint) error {
	_, err := ths.db.Queryx(
		`UPDATE "storyroom".posts SET is_deleted = true, updated_at = now() WHERE id = $1`,
		postID,
	)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		log.Printf("delete post: %s", err.Error())
		return err
	}

	return nil
}
