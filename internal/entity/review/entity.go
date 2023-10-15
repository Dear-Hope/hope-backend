package review

import (
	"database/sql"
	"fmt"
	"time"
)

type (
	Review struct {
		Id        uint64         `db:"id"`
		ExpertId  uint64         `db:"expert_id"`
		Rating    int64          `db:"rating"`
		Review    sql.NullString `db:"review"`
		CreatedAt time.Time      `db:"created_at"`
		UpdatedAt sql.NullTime   `db:"updated_at"`
		IsDeleted bool           `db:"is_deleted"`
	}

	Reviews []Review
)

func (ths Review) ToResponse() *Response {
	return &Response{
		Id:        ths.Id,
		ExpertId:  ths.ExpertId,
		Rating:    ths.Rating,
		Review:    ths.Review.String,
		CreatedAt: ths.CreatedAt.UTC().UnixMilli(),
	}
}

func (ths Reviews) ToListResponse() *ListResponse {
	var (
		sum, totalReview int64
		totalRating      = int64(len(ths))
		res              = make([]Response, totalRating)
		listResponse     = &ListResponse{
			AvgRating: "belum ada penilaian",
		}
	)
	for i, review := range ths {
		if review.Review.Valid {
			totalReview++
		}
		constructTotalStars(review.Rating, listResponse)
		sum += review.Rating
		res[i] = *review.ToResponse()
	}

	if totalRating > 0 {
		listResponse.AvgRating = fmt.Sprintf("%.1f/5.0", float64(sum)/float64(totalRating))
	}

	listResponse.TotalRating = fmt.Sprintf("%d rating", totalRating)
	listResponse.TotalReview = fmt.Sprintf("%d ulasan", totalReview)
	listResponse.Reviews = res

	return listResponse
}

func constructTotalStars(rating int64, listResponse *ListResponse) {
	switch rating {
	case 1:
		listResponse.OneStar++
	case 2:
		listResponse.TwoStar++
	case 3:
		listResponse.ThreeStar++
	case 4:
		listResponse.FourStar++
	case 5:
		listResponse.FiveStar++
	default:
		return
	}
}

type (
	Response struct {
		Id        uint64 `json:"id"`
		ExpertId  uint64 `json:"expertId"`
		Rating    int64  `json:"rating"`
		Review    string `json:"review,omitempty"`
		CreatedAt int64  `json:"createdAt"`
	}

	ListResponse struct {
		AvgRating   string     `json:"avgRating"`
		TotalRating string     `json:"totalRating"`
		TotalReview string     `json:"totalReview"`
		OneStar     int64      `json:"oneStar"`
		TwoStar     int64      `json:"twoStar"`
		ThreeStar   int64      `json:"threeStar"`
		FourStar    int64      `json:"fourStar"`
		FiveStar    int64      `json:"fiveStar"`
		Reviews     []Response `json:"reviews"`
	}
)
