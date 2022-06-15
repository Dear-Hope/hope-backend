package models

type Subscription struct {
	ID           string `db:"id"`
	Email        string `db:"email"`
	SubscribedAt int64  `db:"subscribed_at"`
}

type NewsletterService interface {
	Subscribe(NewSubscriberRequest) error
	Unsubscribe(string) error
}

type NewsletterRepository interface {
	Create(Subscription) error
	Delete(string) error
}

type NewSubscriberRequest struct {
	Email string `json:"email"`
	Time  int64  `json:"time"`
}
