package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetAllCommentByPostID(postID uint) (model.Comments, error) {
	var comments model.Comments

	if postID <= 0 {
		return comments, nil
	}

	err := ths.db.Select(
		&comments,
		`SELECT c.id, c.content, c.created_at, u.id AS author_id, u.alias AS author_name, pr.photo AS author_pic_url 
		FROM "storyroom".comments c, "auth".users u, "auth".profiles pr 
		WHERE c.post_id = $1 AND c.author_id = u.id AND pr.user_id = u.id`,
		postID,
	)
	if err != nil {
		log.Printf("get all comment by post id: %s", err.Error())
		return nil, err
	}

	return comments, nil
}
