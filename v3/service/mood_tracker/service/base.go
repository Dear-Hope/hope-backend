package service

import (
	"HOPE-backend/v3/service/auth"
	"HOPE-backend/v3/service/mood_tracker"
)

type service struct {
	moodRepo mood_tracker.Repository
	userRepo auth.Repository
}

func NewService(
	moodRepo mood_tracker.Repository,
	userRepo auth.Repository,
) mood_tracker.Service {
	return &service{
		moodRepo: moodRepo,
		userRepo: userRepo,
	}
}

// func (ths *service) ListEmotionPerWeek(userID uint) ([]*models.Emotion, error) {
// 	emotions, err := ths.moodRepo.GetAllEmotionByUserIDPerWeek(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return emotions, err
// }

// func (ths *service) ListEmotionPerMonth(userID uint) ([]*models.Emotion, error) {
// 	emotions, err := ths.moodRepo.GetAllEmotionByUserIDPerMonth(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return emotions, err
// }
