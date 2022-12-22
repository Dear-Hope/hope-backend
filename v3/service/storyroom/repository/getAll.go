package repository

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/storyroom/filter"
	"log"
)

func (ths *repository) GetAll(f filter.List) (model.Posts, error) {
	var (
		posts     model.Posts
		args      []interface{}
		sort      = "created_at"
		direction = "DESC"
	)

	from := ` FROM "storyroom".posts p, "auth".users u, "auth".profiles pr`
	where := ` WHERE p.is_deleted = false AND p.author_id = u.id AND pr.user_id = u.id`

	if f.CategoryID > 0 {
		from += `, "storyroom".category_posts cp`
		where += ` AND cp.post_id = p.id AND cp.category_id = $1`
		args = append(args, f.CategoryID)
	}

	if f.UserID > 0 {
		where += ` AND p.author_id = $1`
		args = append(args, f.UserID)
	}

	if f.Sort != "" {
		sort = f.Sort
	}

	if f.Direction != "" {
		direction = f.Direction
	}

	query := `SELECT p.id, p.content, p.created_at, u.id AS author_id, u.alias AS author_name, pr.photo AS author_pic_url` + from + where

	err := ths.db.Select(
		&posts,
		query+` ORDER BY p.`+sort+` `+direction,
		args...,
	)
	if err != nil {
		log.Printf("get all post: %s", err.Error())
		return nil, err
	}

	return posts, nil
}
