package psychologicalrecord

import (
	"HOPE-backend/v1/models"
	"errors"
	"time"
)

type service struct {
	recordRepo models.PsychologicalRecordRepository
	userRepo   models.AuthRepository
}

func NewPsychologicalRecordService(
	recordRepo models.PsychologicalRecordRepository,
	userRepo models.AuthRepository,
) models.PsychologicalRecordService {
	return &service{
		recordRepo: recordRepo,
		userRepo:   userRepo,
	}
}

func (ths *service) NewRecord(req models.NewPsychologicalRecordRequest, psychologistID uint) (
	*models.PsychologicalRecord,
	error,
) {
	psychologist, err := ths.userRepo.GetUserByID(psychologistID)
	if err != nil {
		return nil, errors.New("psychologist with given ID not found")
	}

	if psychologist.Role != "psychologist" {
		return nil, errors.New("the one who filled this record was not a psyhologist")
	}

	patient, err := ths.userRepo.GetUserByID(req.PatientID)
	if err != nil {
		return nil, errors.New("patient with given ID not found")
	}

	if patient.Role != "patient" {
		return nil, errors.New("the one being filled in this record was not a patient")
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.New("failed to parse date given")
	}

	newRecord := models.PsychologicalRecord{
		PsychologistName:            psychologist.FirstName + " " + psychologist.LastName,
		PatientName:                 req.PatientName,
		Date:                        date,
		Sex:                         req.Sex,
		Age:                         req.Age,
		Education:                   req.Education,
		Occupation:                  req.Occupation,
		JobPosition:                 req.JobPosition,
		Habit:                       req.Habit,
		Analysis:                    req.Analysis,
		Problem:                     req.Problem,
		Recommendation:              req.Recommendation,
		Note:                        req.Note,
		PsychologistAcknowledgement: true,
		Psychologist:                *psychologist,
		Patient:                     *patient,
	}

	record, err := ths.recordRepo.Create(newRecord)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (ths *service) GetRecord(id uint) (*models.PsychologicalRecord, error) {
	record, err := ths.recordRepo.GetRecordByID(id)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (ths *service) ListRecord(psychologistID uint) ([]*models.PsychologicalRecord, error) {
	records, err := ths.recordRepo.GetAllRecordByPyschologistID(psychologistID)
	if err != nil {
		return nil, err
	}

	return records, err
}
