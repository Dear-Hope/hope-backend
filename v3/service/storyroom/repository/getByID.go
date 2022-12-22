package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetByID(id, userID uint) (*model.Post, error) {
	var post model.Post

	err := ths.db.Get(
		&post,
		`SELECT p.id, p.content, p.created_at, u.id AS author_id, u.alias AS author_name, pr.photo AS author_pic_url 
		FROM `+post.TableWithSchemaName()+` p, "auth".users u, "auth".profiles pr 
		WHERE p.is_deleted = false AND p.id = $1 AND p.author_id = u.id AND pr.user_id = u.id`,
		id,
	)
	if err != nil {
		log.Printf("get detail post: %s", err.Error())
		return nil, err
	}

	return &post, nil
}
