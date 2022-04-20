package models

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model
	Email       string `json:"email" gorm:"not null;unique"`
	SubsribedAt int64  `json:"subscribed_at" gorm:"not null"`
	MemberID    string `json:"member_id" gorm:"not null"`
}

func (Subscription) TableName() string {
	return "newsletter"
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
