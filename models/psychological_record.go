package models

import (
	"time"

	"gorm.io/gorm"
)

type PsychologicalRecord struct {
	gorm.Model
	PsychologistName            string    `json:"psychologist_name" gorm:"not null"`
	PatientName                 string    `json:"patient_name" gorm:"not null"`
	Date                        time.Time `json:"date" gorm:"not null"`
	Sex                         string    `json:"sex" gorm:"not null;not blank"`
	Age                         int       `json:"age" gorm:"not null;not blank"`
	Education                   string    `json:"education"`
	Occupation                  string    `json:"occupation"`
	JobPosition                 string    `json:"job_position"`
	Habit                       string    `json:"habit"`
	Problem                     string    `json:"problem"`
	Analysis                    string    `json:"analysis"`
	Recommendation              string    `json:"recommendation"`
	Note                        string    `json:"note"`
	PsychologistAcknowledgement bool      `json:"psychologist_acknowledgement" gorm:"not null"`
	PsychologistID              uint      `json:"psychologist_id" gorm:"uniqueIndex:idx_psychologist_patient_id"`
	PatientID                   uint      `json:"patient_id" gorm:"uniqueIndex:idx_psychologist_patient_id"`
	Psychologist                User      `json:"-" gorm:"constraint:OnUpdate:CASCADE;"`
	Patient                     User      `json:"-" gorm:"constraint:OnUpdate:CASCADE;"`
}

func (PsychologicalRecord) TableName() string {
	return "psychological_record"
}

type PsychologicalRecordService interface {
	NewRecord(PsychologicalRecord) (*PsychologicalRecord, error)
	GetRecord(uint) (*PsychologicalRecord, error)
	ListRecord(uint) ([]*PsychologicalRecord, error)
}

type PsychologicalRecordRepository interface {
	Create(PsychologicalRecord) (*PsychologicalRecord, error)
	GetRecordByID(uint) (*PsychologicalRecord, error)
	GetAllRecordByPyschologistID(uint) ([]*PsychologicalRecord, error)
}
