package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) StoreComment(newComment model.Comment) (*model.Comment, error) {
	rows, err := ths.db.NamedQuery(
		`WITH rows AS (
			INSERT INTO "storyroom".comments (post_id, author_id, content) 
			VALUES (:post_id, :author_id, :content) 
			RETURNING id, author_id, created_at
		)
		SELECT r.id, r.created_at, u.alias AS author_name, pr.photo AS author_pic_url 
		FROM rows r, "auth".users u, "auth".profiles pr
		WHERE r.author_id = u.id AND pr.user_id = u.id
		`,
		newComment,
	)
	if err != nil {
		log.Printf("new post comment failed: %s", err.Error())
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		if err = rows.Scan(&newComment.ID, &newComment.CreatedAt, &newComment.AuthorName, &newComment.AuthorPicURL); err != nil {
			log.Printf("new post comment failed: %s", err.Error())
			return nil, err
		}
	}

	return &newComment, nil
}
