package repository

import (
	"log"
)

func (ths *repository) UpsertLike(postID, userID uint) error {
	_, err := ths.db.NamedQuery(
		`INSERT INTO "storyroom".likes (post_id, user_id) 
			VALUES (:post_id, :user_id) 
			ON CONFLICT (post_id, user_id)
			DO UPDATE SET updated_at = now(), is_deleted = NOT "storyroom".likes.is_deleted`,
		map[string]interface{}{"post_id": postID, "user_id": userID},
	)
	if err != nil {
		log.Printf("upsert like failed: %s", err.Error())
		return err
	}

	return nil
}
