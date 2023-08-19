package expert

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
		Id          uint64         `db:"id"`
		Email       string         `db:"email"`
		Password    string         `db:"password"`
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

func (ths Expert) ToResponse() *Response {
	return &Response{
		Id:          ths.Id,
		Email:       ths.Email,
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

func (ths Experts) ToListResponse() []Response {
	res := make([]Response, len(ths))
	for i, expert := range ths {
		res[i] = *expert.ToResponse()
	}

	return res
}

type (
	Response struct {
		Id          uint64   `json:"id"`
		Email       string   `json:"email,omitempty"`
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
	CreateUpdateRequest struct {
		Id         uint64   `json:"id,omitempty"`
		Email      string   `json:"email"`
		Password   string   `json:"password"`
		Name       string   `json:"name"`
		Expertise  string   `json:"expertise"`
		Rating     float32  `json:"rating,omitempty"`
		Price      int      `json:"price"`
		Title      string   `json:"title"`
		Education  string   `json:"education"`
		Experience string   `json:"experience"`
		Photo      string   `json:"photo,omitempty"`
		Bio        string   `json:"bio,omitempty"`
		TopicIds   []uint64 `json:"topicIds"`
	}
)

func (ths CreateUpdateRequest) ToExpertModel() Expert {
	var photo, bio sql.NullString
	if ths.Photo != "" {
		photo = sql.NullString{Valid: true, String: ths.Photo}
	}

	if ths.Bio != "" {
		bio = sql.NullString{Valid: true, String: ths.Bio}
	}

	return Expert{
		Id:         ths.Id,
		Name:       ths.Name,
		Email:      ths.Email,
		Password:   ths.Password,
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
