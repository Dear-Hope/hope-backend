package repository

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/counseling/filter"
	"log"
	"strings"
)

func (ths *repository) GetAllExperts(f filter.ListExpert) (model.Experts, error) {
	var (
		args    []interface{}
		experts model.Experts
	)
	query := `WITH row AS (
		SELECT et.expert_id, array_agg(t.name) as topic_names 
		FROM "counseling".expert_topics et, "counseling".topics t 
		WHERE et.topic_id = t.id GROUP BY et.expert_id
	)
	SELECT DISTINCT e.id, e.name, e.photo, e.expertise, e.rating, e.price, e.is_available, r.topic_names 
	FROM "counseling".experts e, row r, "counseling".expert_topics et 
	WHERE e.id = r.expert_id AND e.id = et.expert_id AND e.is_deleted = false `

	if f.TopicId > 0 {
		query += "AND et.topic_id = ? "
		args = append(args, f.TopicId)
	}

	if f.Expertise != "" {
		query += "AND e.expertise = ? "
		args = append(args, strings.ToUpper(f.Expertise))
	}

	err := ths.db.Select(
		&experts,
		ths.db.Rebind(query+"ORDER BY e.is_available DESC"),
		args...,
	)
	if err != nil {
		log.Printf("get all experts: %s", err.Error())
		return nil, err
	}

	return experts, nil
}
