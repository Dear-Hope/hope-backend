package repository

import (
	"database/sql"
	"errors"
	"log"
)

func (ths *repository) DeleteCommentByPostID(postID uint) error {
	_, err := ths.db.Queryx(
		`UPDATE "storyroom".comments SET is_deleted = true, updated_at = now() WHERE post_id = $1`,
		postID,
	)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		log.Printf("delete comment: %s", err.Error())
		return err
	}

	return nil
}
