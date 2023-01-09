package repository

import (
	"HOPE-backend/v3/model"
	"fmt"
	"log"
	"strings"
)

func (ths *repository) GetAllCommentByPostID(postID uint, excludedIDs []uint) (model.Comments, error) {
	var (
		comments model.Comments
		where    string
		args     = []interface{}{postID}
	)

	if postID <= 0 {
		return comments, nil
	}

	if len(excludedIDs) > 0 {
		var ids []string
		for _, id := range excludedIDs {
			ids = append(ids, fmt.Sprintf("%d", id))
		}

		where = fmt.Sprintf(` AND c.author_id NOT IN (%s)`, strings.Join(ids, ","))
	}

	err := ths.db.Select(
		&comments,
		`SELECT c.id, c.content, c.created_at, u.id AS author_id, u.alias AS author_name, pr.photo AS author_pic_url 
		FROM "storyroom".comments c, "auth".users u, "auth".profiles pr 
		WHERE c.post_id = $1 AND c.author_id = u.id AND pr.user_id = u.id`+where,
		args...,
	)
	if err != nil {
		log.Printf("get all comment by post id: %s", err.Error())
		return nil, err
	}

	return comments, nil
}
