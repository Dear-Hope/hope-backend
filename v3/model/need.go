package model

type (
	Need struct {
		ID   uint   `db:"id"`
		Name string `db:"name"`
	}

	Needs []Need
)

func (ths Need) TableWithSchemaName() string {
	return `"selfcare".needs`
}

type (
	NeedResponse struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
)
