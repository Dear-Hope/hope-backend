package models

type SelfCareItem struct {
	ID          uint   `json:"id" db:"id"`
	Link        string `json:"link" db:"link"`
	Mood        string `json:"mood" db:"mood"`
	Title       string `json:"title" db:"title"`
	Type        string `json:"type" db:"type"`
	Description string `json:"description" db:"description"`
}

type SelfCareService interface {
	NewItem(NewSelfCareItemRequest) (*SelfCareItem, error)
	ListItems() ([]*SelfCareItem, error)
	GetItemsByMood(string) ([]*SelfCareItem, error)
}

type SelfCareRepository interface {
	Create(SelfCareItem) (*SelfCareItem, error)
	GetAllItems() ([]*SelfCareItem, error)
	GetItemsByMood(string) ([]*SelfCareItem, error)
}

type NewSelfCareItemRequest struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Mood        string `json:"mood"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}
