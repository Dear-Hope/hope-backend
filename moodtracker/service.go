package moodtracker

import (
	"HOPE-backend/models"
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

func (ths *service) NewEmotion(req models.NewEmotionRequest, patientID uint) (
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

	patient, err := ths.userRepo.GetUserByID(patientID)
	if err != nil {
		return nil, errors.New("patient with given ID not found")
	}

	if patient.Role != "patient" {
		return nil, errors.New("the one who filled this record was not a patient")
	}

	newEmotion := models.Emotion{
		Mood:        req.Mood,
		Triggers:    req.Triggers,
		Description: req.Description,
		TimeFrame:   timeFrame,
		PatientID:   patientID,
		Patient:     *patient,
	}

	emotion, err := ths.moodRepo.Create(newEmotion)
	if err != nil {
		return nil, err
	}

	return emotion, nil
}

func (ths *service) ListEmotion(patientID uint) ([]*models.Emotion, error) {
	emotions, err := ths.moodRepo.GetAllEmotionByPatientID(patientID)
	if err != nil {
		return nil, err
	}

	return emotions, err
}

func (ths *service) ListEmotionPerWeek(patientID uint) ([]*models.Emotion, error) {
	emotions, err := ths.moodRepo.GetAllEmotionByPatientIDPerWeek(patientID)
	if err != nil {
		return nil, err
	}

	return emotions, err
}

func (ths *service) ListEmotionPerMonth(patientID uint) ([]*models.Emotion, error) {
	emotions, err := ths.moodRepo.GetAllEmotionByPatientIDPerMonth(patientID)
	if err != nil {
		return nil, err
	}

	return emotions, err
}
