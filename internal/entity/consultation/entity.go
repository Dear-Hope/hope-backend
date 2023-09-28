package consultation

import (
	"database/sql"
	"strings"
	"time"
)

type Consultation struct {
	Id          uint64       `db:"id"`
	UserId      uint64       `db:"user_id"`
	ExpertId    uint64       `db:"expert_id"`
	TypeId      uint64       `db:"type_id"`
	BookingDate string       `db:"booking_date"`
	StartTime   string       `db:"start_time"`
	EndTime     string       `db:"end_time"`
	Status      string       `db:"status"`
	UserNotes   string       `db:"user_notes"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
	IsDeleted   bool         `db:"is_deleted"`
}

type Consultations []Consultation

type (
	ExpertListRequest struct {
		UserId       uint64
		ExpertId     uint64
		BookingDate  string
		BookingMonth string
		Status       Status
	}
)

type (
	ExpertResponse struct {
		Id                  uint64 `json:"id,omitempty"`
		ClientName          string `json:"clientName,omitempty"`
		ClientPhoto         string `json:"clientPhoto,omitempty"`
		ClientNote          string `json:"clientNote,omitempty"`
		TypeId              uint64 `json:"typeId,omitempty"`
		Status              string `json:"status,omitempty"`
		Time                string `json:"time,omitempty"`
		IsStartConsultation bool   `json:"isStartConsultation"`
	}

	ExpertListResponse struct {
		Data        []ExpertResponse `json:"data"`
		TotalClient int              `json:"totalClient"`
	}
)

type Status int

var (
	status     = []string{"SCHEDULED", "ACCEPTED", "REJECTED", "ONGOING", "COMPLETED"}
	statusText = []string{"Klien Baru", "Klien", "Ditolak", "Berlangsung", "Selesai"}
)

const (
	Scheduled Status = iota + 1
	Accepted
	Rejected
	Ongoing
	Completed
)

func (s Status) String() string {
	return status[s-1]
}

func (s Status) Text() string {
	return statusText[s-1]
}

func GetStatus(s string) Status {
	for idx := range status {
		if strings.EqualFold(status[idx], s) {
			return Status(idx + 1)
		}
	}
	return 1
}
