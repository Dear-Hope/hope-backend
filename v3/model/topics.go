package model

type (
	Topic struct {
		ID       uint   `db:"id"`
		Name     string `db:"name"`
		ImageURL string `db:"image_url"`
	}

	Topics []Topic
)

func (ths Topic) ToTopicResponse() *TopicResponse {
	return &TopicResponse{
		ID:       ths.ID,
		Name:     ths.Name,
		ImageURL: ths.ImageURL,
	}
}

func (ths Topics) ToListTopicResponse() []TopicResponse {
	res := make([]TopicResponse, len(ths))
	for i, topic := range ths {
		res[i] = *topic.ToTopicResponse()
	}

	return res
}

type (
	TopicResponse struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		ImageURL string `json:"imageUrl"`
	}
)
