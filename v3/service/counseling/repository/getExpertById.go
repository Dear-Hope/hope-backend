package repository

import (
	"HOPE-backend/v3/model"
	"log"
)

func (ths *repository) GetExpertById(id uint) (*model.Expert, error) {
	var expert model.Expert
	query := `SELECT e.id, e.name, e.photo, e.expertise, e.rating, e.price, e.is_available, e.title, 
       		e.bio, e.education, e.experience, array_agg(t.name) as topic_names
			FROM "counseling".experts e, "counseling".topics t, "counseling".expert_topics et
			WHERE e.id = et.expert_id AND t.id = et.topic_id AND e.id = ? GROUP BY e.id`

	err := ths.db.Get(&expert, ths.db.Rebind(query), id)
	if err != nil {
		log.Printf("get expert by id: %s", err.Error())
		return nil, err
	}

	return &expert, nil
}
