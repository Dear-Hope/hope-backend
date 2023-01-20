package model

import (
	"database/sql"
	"github.com/lib/pq"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"time"
)

type (
	Expert struct {
		ID          uint           `db:"id"`
		Name        string         `db:"name"`
		Expertise   string         `db:"expertise"`
		Rating      float32        `db:"rating"`
		Price       int            `db:"price"`
		IsAvailable bool           `db:"is_available"`
		Title       string         `db:"title"`
		Education   string         `db:"education"`
		Experience  string         `db:"experience"`
		Photo       sql.NullString `db:"photo"`
		Bio         sql.NullString `db:"bio"`
		Topics      pq.StringArray `db:"topic_names"`
		CreatedAt   time.Time      `db:"created_at"`
		UpdatedAt   time.Time      `db:"updated_at"`
		IsDeleted   bool           `db:"is_deleted"`
	}

	Experts []Expert
)

func (ths Expert) ToExpertResponse() *ExpertResponse {
	return &ExpertResponse{
		ID:          ths.ID,
		Name:        ths.Name,
		Expertise:   cases.Title(language.Und).String(ths.Expertise),
		Rating:      ths.Rating,
		Price:       ths.Price,
		IsAvailable: ths.IsAvailable,
		Title:       ths.Title,
		Education:   ths.Education,
		Experience:  ths.Experience,
		Photo:       ths.Photo.String,
		Bio:         ths.Bio.String,
		Topics:      ths.Topics,
	}
}

func (ths Experts) ToListExpertResponse() []ExpertResponse {
	res := make([]ExpertResponse, len(ths))
	for i, expert := range ths {
		res[i] = *expert.ToExpertResponse()
	}

	return res
}

type (
	ExpertResponse struct {
		ID          uint     `json:"id"`
		Name        string   `json:"name,omitempty"`
		Expertise   string   `json:"expertise,omitempty"`
		Rating      float32  `json:"rating,omitempty"`
		Price       int      `json:"price,omitempty"`
		IsAvailable bool     `json:"isAvailable"`
		Title       string   `json:"title,omitempty"`
		Education   string   `json:"education,omitempty"`
		Experience  string   `json:"experience,omitempty"`
		Photo       string   `json:"photo,omitempty"`
		Bio         string   `json:"bio,omitempty"`
		Topics      []string `json:"topics"`
	}
)

type (
	CreateUpdateExpertRequest struct {
		ID         uint    `json:"id,omitempty"`
		Name       string  `json:"name"`
		Expertise  string  `json:"expertise"`
		Rating     float32 `json:"rating,omitempty"`
		Price      int     `json:"price"`
		Title      string  `json:"title"`
		Education  string  `json:"education"`
		Experience string  `json:"experience"`
		Photo      string  `json:"photo,omitempty"`
		Bio        string  `json:"bio,omitempty"`
		TopicIds   []uint  `json:"topicIds"`
	}
)

func (ths CreateUpdateExpertRequest) ToExpertModel() Expert {
	var photo, bio sql.NullString
	if ths.Photo != "" {
		photo = sql.NullString{Valid: true, String: ths.Photo}
	}

	if ths.Bio != "" {
		bio = sql.NullString{Valid: true, String: ths.Bio}
	}

	return Expert{
		ID:         ths.ID,
		Name:       ths.Name,
		Expertise:  strings.ToUpper(ths.Expertise),
		Rating:     ths.Rating,
		Price:      ths.Price,
		Title:      ths.Title,
		Education:  ths.Education,
		Experience: ths.Experience,
		Photo:      photo,
		Bio:        bio,
	}
}
