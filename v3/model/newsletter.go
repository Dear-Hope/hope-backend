package model

type Subscription struct {
	ID           string `db:"id"`
	Email        string `db:"email"`
	SubscribedAt int64  `db:"subscribed_at"`
}

type NewSubscriberRequest struct {
	Email string `json:"email"`
	Time  int64  `json:"time"`
}
