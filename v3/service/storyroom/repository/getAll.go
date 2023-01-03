package repository

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/storyroom/filter"
	"fmt"
	"log"
	"strings"
)

func (ths *repository) GetAll(f filter.List) (model.Posts, error) {
	var (
		posts       model.Posts
		args        []interface{}
		sort        = "created_at"
		direction   = "DESC"
		argsCounter = 1
	)

	from := ` FROM "storyroom".posts p, "auth".users u, "auth".profiles pr`
	where := ` WHERE p.is_deleted = false AND p.author_id = u.id AND pr.user_id = u.id`

	if f.CategoryID > 0 {
		from += `, "storyroom".category_posts cp`
		where += fmt.Sprintf(` AND cp.post_id = p.id AND cp.category_id = $%d`, argsCounter)
		args = append(args, f.CategoryID)
		argsCounter++
	}

	if f.UserID > 0 {
		where += fmt.Sprintf(` AND p.author_id = $%d`, argsCounter)
		args = append(args, f.UserID)
		argsCounter++
	}

	if f.Sort != "" {
		sort = f.Sort
	}

	if f.Direction != "" {
		direction = f.Direction
	}

	if len(f.ExcludedID) > 0 {
		var ids []string
		for _, id := range f.ExcludedID {
			ids = append(ids, fmt.Sprintf("%d", id))
		}

		where += fmt.Sprintf(` AND p.author_id NOT IN (%s)`, strings.Join(ids, ","))
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
