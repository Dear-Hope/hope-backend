package model

type (
	Category struct {
		ID          uint   `db:"id"`
		Name        string `db:"name"`
		ImageUrl    string `db:"image_url"`
		Description string `db:"description"`
	}

	Categories []Category
)

func (ths Category) TableWithSchemaName() string {
	return `"selfcare".categories`
}

func (ths Category) ToCategoryResponse(total int) *CategoryResponse {
	return &CategoryResponse{
		ID:          ths.ID,
		Name:        ths.Name,
		ImageUrl:    ths.ImageUrl,
		Description: ths.Description,
		Total:       total,
	}
}

type (
	CategoryResponse struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Total       int    `json:"total"`
		ImageUrl    string `json:"imageUrl"`
		Description string `json:"description"`
	}
)
