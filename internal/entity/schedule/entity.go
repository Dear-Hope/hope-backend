package schedule

import (
	"github.com/lib/pq"
	"time"
)

var Days = []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

type (
	DbSchedule struct {
		Schedule
		Timeslot
	}

	DbSchedules []DbSchedule

	Schedule struct {
		Id        uint64    `db:"id"`
		Day       string    `db:"day"`
		IsActive  bool      `db:"is_active"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		IsDeleted bool      `db:"is_deleted"`
	}

	Schedules []Schedule

	Timeslot struct {
		Id         uint64        `db:"id"`
		ScheduleId uint64        `db:"schedule_id"`
		StartTime  string        `db:"start_time"`
		EndTime    string        `db:"end_time"`
		TypeIds    pq.Int64Array `db:"type_ids"`
		CreatedAt  time.Time     `db:"created_at"`
		UpdatedAt  time.Time     `db:"updated_at"`
		IsDeleted  bool          `db:"is_deleted"`
	}

	Timeslots []Timeslot
)

type (
	Response struct {
		Id        uint64             `json:"id"`
		Day       string             `json:"day"`
		IsActive  bool               `json:"isActive"`
		Timeslots []TimeslotResponse `json:"timeslots,omitempty"`
	}

	TimeslotResponse struct {
		Id        uint64  `json:"id,omitempty"`
		StartTime string  `json:"startTime,omitempty"`
		EndTime   string  `json:"endTime,omitempty"`
		TypeIds   []int64 `json:"typeIds,omitempty"`
	}

	TimeslotUserResponse struct {
		StartTime string `json:"startTime,omitempty"`
		EndTime   string `json:"endTime,omitempty"`
		IsBooked  bool   `json:"isBooked"`
	}
)

type (
	UpdateRequest struct {
		ExpertId  uint64                  `json:"expertId"`
		Schedules []UpdateScheduleRequest `json:"schedules"`
	}

	UpdateScheduleRequest struct {
		Id        uint64                 `json:"id"`
		IsActive  bool                   `json:"isActive"`
		Timeslots UpdateTimeslotRequests `json:"timeslots"`
	}

	UpdateTimeslotRequest struct {
		Id        uint64  `json:"id"`
		StartTime string  `json:"startTime"`
		EndTime   string  `json:"endTime"`
		TypeIds   []int64 `json:"typeIds"`
	}

	UpdateTimeslotRequests []UpdateTimeslotRequest
)

func (tr UpdateTimeslotRequest) ToTimeslot(scheduleId uint64) Timeslot {
	return Timeslot{
		Id:         tr.Id,
		ScheduleId: scheduleId,
		StartTime:  tr.StartTime,
		EndTime:    tr.EndTime,
		TypeIds:    tr.TypeIds,
	}
}

func (trs UpdateTimeslotRequests) ToTimeslots(scheduleId uint64) Timeslots {
	response := make(Timeslots, len(trs))
	for i, tr := range trs {
		response[i] = tr.ToTimeslot(scheduleId)
	}

	return response
}

var (
	DayIndonesian    = []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
	MonthsIndonesian = []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September",
		"Oktober", "November", "Desember"}
)

//func (ts *Timeslots) Scan(src interface{}) error {
//	return pq.GenericArray{A: ts}.Scan(src)
//}
//
//func (t *Timeslot) Scan(src interface{}) error {
//	var data []byte
//	switch v := src.(type) {
//	case string:
//		data = []byte(v)
//	case []byte:
//		data = v
//	}
//	return json.Unmarshal(data, t)
//}
