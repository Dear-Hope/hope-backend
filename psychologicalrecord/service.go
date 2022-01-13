package psychologicalrecord

import (
	"HOPE-backend/models"
	"errors"
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

func (ths *service) NewRecord(newRecord models.PsychologicalRecord) (*models.PsychologicalRecord, error) {
	psychologist, err := ths.userRepo.GetUserByID(newRecord.PsychologistID)
	if err != nil {
		return nil, errors.New("psychologist with given ID not found")
	}

	if psychologist.Role != "psychologist" {
		return nil, errors.New("the one who filled this record was not a psyhologist")
	}

	patient, err := ths.userRepo.GetUserByID(newRecord.PatientID)
	if err != nil {
		return nil, errors.New("patient with given ID not found")
	}

	if patient.Role != "patient" {
		return nil, errors.New("the one being filled in this record was not a patient")
	}

	newRecord.Psychologist = *psychologist
	newRecord.Patient = *patient
	newRecord.PsychologistAcknowledgement = true

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
