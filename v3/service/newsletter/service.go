package newsletter

import "HOPE-backend/v3/model"

type Service interface {
	Subscribe(model.NewSubscriberRequest) *model.ServiceError
	Unsubscribe(string) *model.ServiceError
}
