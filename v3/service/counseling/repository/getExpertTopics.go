package repository

import "log"

func (ths *repository) GetExpertTopics(expertId uint) ([]uint, error) {
	var topicIds []uint
	err := ths.db.Select(
		&topicIds,
		ths.db.Rebind(`SELECT topic_id FROM "counseling".expert_topics WHERE expert_id = ?`),
		expertId,
	)
	if err != nil {
		log.Printf("get expert's topics: %s", err.Error())
		return nil, err
	}

	return topicIds, nil
}
