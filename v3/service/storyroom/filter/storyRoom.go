package filter

type List struct {
	CategoryID uint
	Sort       string
	Direction  string
	UserID     uint
	ExcludedID []uint
}
