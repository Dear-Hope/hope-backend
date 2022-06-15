package moodtracker

import (
	"HOPE-backend/v2/models"
	"errors"
)

type service struct {
	moodRepo models.MoodTrackerRepository
	userRepo models.AuthRepository
}

func NewMoodTrackerService(
	moodRepo models.MoodTrackerRepository,
	userRepo models.AuthRepository,
) models.MoodTrackerService {
	return &service{
		moodRepo: moodRepo,
		userRepo: userRepo,
	}
}

func (ths *service) NewEmotion(req models.NewEmotionRequest, userID uint) (
	*models.Emotion,
	error,
) {
	if !req.IsMoodAvailable() {
		return nil, errors.New("mood given is not listed in our database")
	}

	timeFrame, err := req.ConvertIntoTimeFrame()
	if err != nil {
		return nil, err
	}

	newEmotion := models.Emotion{
		Mood:        req.Mood,
		Triggers:    req.Triggers,
		Description: req.Description,
		TimeFrame:   timeFrame,
		Date:        req.Time,
		UserID:      userID,
	}

	emotion, err := ths.moodRepo.Create(newEmotion)
	if err != nil {
		return nil, err
	}

	return emotion, nil
}

func (ths *service) ListEmotion(userID uint) ([]*models.Emotion, error) {
	emotions, err := ths.moodRepo.GetAllEmotionByUserID(userID)
	if err != nil {
		return nil, err
	}

	return emotions, err
}

func (ths *service) ListEmotionPerWeek(userID uint) ([]*models.Emotion, error) {
	emotions, err := ths.moodRepo.GetAllEmotionByUserIDPerWeek(userID)
	if err != nil {
		return nil, err
	}

	return emotions, err
}

func (ths *service) ListEmotionPerMonth(userID uint) ([]*models.Emotion, error) {
	emotions, err := ths.moodRepo.GetAllEmotionByUserIDPerMonth(userID)
	if err != nil {
		return nil, err
	}

	return emotions, err
}
