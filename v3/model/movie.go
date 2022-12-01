package model

import "github.com/lib/pq"

type (
	Movie struct {
		ID          uint           `db:"id"`
		Title       string         `db:"title"`
		Year        int            `db:"year"`
		Country     string         `db:"country"`
		Genres      pq.StringArray `db:"genres"`
		Description string         `db:"description"`
		TrailerLink string         `db:"trailer_link"`
		PosterLink  string         `db:"poster_link"`

		Moods Moods
		Needs Needs
	}

	Movies []Movie
)

func (ths Movie) TableWithSchemaName() string {
	return `"selfcare".movies`
}

func (ths Movie) ToMovieResponse() *MovieResponse {
	var moods, needs []string

	if ths.Moods != nil {
		for _, mood := range ths.Moods {
			moods = append(moods, mood.Name)
		}
	}

	if ths.Needs != nil {
		for _, need := range ths.Needs {
			needs = append(needs, need.Name)
		}
	}

	return &MovieResponse{
		ID:          ths.ID,
		Title:       ths.Title,
		Year:        ths.Year,
		Country:     ths.Country,
		Genres:      ths.Genres,
		Description: ths.Description,
		TrailerLink: ths.TrailerLink,
		PosterLink:  ths.PosterLink,
		Moods:       moods,
		Needs:       needs,
	}
}

func (ths Movies) ToListMovieResponse() []MovieResponse {
	res := make([]MovieResponse, len(ths))
	for i, movie := range ths {
		res[i] = *movie.ToMovieResponse()
	}

	return res
}

type (
	MovieResponse struct {
		ID          uint     `json:"id"`
		Title       string   `json:"title"`
		Year        int      `json:"year"`
		Country     string   `json:"country"`
		Genres      []string `json:"genres"`
		Description string   `json:"description"`
		TrailerLink string   `json:"trailer_link"`
		PosterLink  string   `json:"poster_link"`
		Moods       []string `json:"moods"`
		Needs       []string `json:"needs"`
	}

	MovieRequest struct {
		Title       string   `json:"title"`
		Year        int      `json:"year"`
		Country     string   `json:"country"`
		Genres      []string `json:"genres"`
		Description string   `json:"description"`
		TrailerLink string   `json:"trailer_link"`
		PosterLink  string   `json:"poster_link"`
		MoodIDs     []uint   `json:"moodIds"`
		NeedIDs     []uint   `json:"needIds"`
	}
)
