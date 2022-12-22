package repository

import (
	"database/sql"
	"errors"
	"log"
)

func (ths *repository) DeletePostCategoryByPostID(postID uint) error {
	_, err := ths.db.Queryx(
		`DELETE FROM "storyroom".category_posts WHERE post_id = $1`,
		postID,
	)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		log.Printf("delete category post: %s", err.Error())
		return err
	}

	return nil
}
