package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetAllCategoryByPostID(postID uint) (model.PostCategories, error) {
	var categories model.PostCategories

	if postID <= 0 {
		return categories, nil
	}

	err := ths.db.Select(
		&categories,
		`SELECT c.id, c.name 
		FROM "storyroom".category_posts cp, "storyroom".categories c 
		WHERE cp.post_id = $1 AND cp.category_id = c.id`,
		postID,
	)
	if err != nil {
		log.Printf("get all categories by post id: %s", err.Error())
		return nil, err
	}

	return categories, nil
}
