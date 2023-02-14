package filter

type ListExpert struct {
	TopicId   uint
	Expertise string
}

type ListExpertSchedule struct {
	ExpertId int64
	TypeId   int64
	Date     string
	Offset   int
}
