package newsletter

import "HOPE-backend/v3/model"

type Repository interface {
	Create(model.Subscription) error
	Delete(string) error
}
