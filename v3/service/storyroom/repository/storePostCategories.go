package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) StorePostCategories(postID uint, categoryIDs []uint) (model.PostCategories, error) {
	var (
		categories = make(model.PostCategories, len(categoryIDs))
	)

	for i, category := range categories {
		rows, err := ths.db.NamedQuery(
			`WITH rows AS (
				INSERT INTO "storyroom".category_posts (post_id, category_id) 
				VALUES (:post_id, :category_id) 
				RETURNING category_id
			)
			SELECT id, name FROM rows, "storyroom".categories WHERE id = rows.category_id`,
			map[string]interface{}{"post_id": postID, "category_id": categoryIDs[i]},
		)
		if err != nil {
			log.Printf("store post categories failed: %s", err.Error())
			return nil, err
		}

		for rows.Next() {
			if err = rows.Scan(&category.ID, &category.Name); err != nil {
				log.Printf("store post categories failed: %s", err.Error())
				return nil, err
			}
		}

		categories[i] = category
		_ = rows.Close()
	}

	return categories, nil
}
