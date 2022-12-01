package model

type (
	Category struct {
		ID   uint   `db:"id"`
		Name string `db:"name"`
	}

	Categories []Category
)

func (ths Category) TableWithSchemaName() string {
	return `"selfcare".categories`
}

func (ths Category) ToCategoryResponse(total int) *CategoryResponse {
	return &CategoryResponse{
		ID:    ths.ID,
		Name:  ths.Name,
		Total: total,
	}
}

type (
	CategoryResponse struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Total int    `json:"total"`
	}
)
