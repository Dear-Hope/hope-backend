package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) StorePost(newPost model.Post) (*model.Post, error) {
	rows, err := ths.db.NamedQuery(
		`WITH rows AS (
			INSERT INTO `+newPost.TableWithSchemaName()+` (author_id, content) 
			VALUES (:author_id, :content) 
			RETURNING id, author_id, created_at
		)
		SELECT r.id, r.created_at, u.alias AS author_name, pr.photo AS author_pic_url 
		FROM rows r, "auth".users u, "auth".profiles pr
		WHERE r.author_id = u.id AND pr.user_id = u.id
		`,
		newPost,
	)
	if err != nil {
		log.Printf("new post create failed: %s", err.Error())
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		if err = rows.Scan(&newPost.ID, &newPost.CreatedAt, &newPost.AuthorName, &newPost.AuthorPicURL); err != nil {
			log.Printf("new post create failed: %s", err.Error())
			return nil, err
		}
	}

	return &newPost, nil
}
