package domain

type StatisticByTypeQuery struct {
	StartDate     string
	EndDate       string
	ArchiveTypeId int
}

//type StatisticWithTagQuery

type TotalListByRootTagQuery struct {
	StartDate     string
	EndDate       string
	ArchiveTypeId int
}

type StatisticWithTag interface {
	Total() int
	TagId() int
	TagName() string
}

type statisticWithTag struct {
	total   int
	tagId   int
	tagName string
}

func NewStatisticWithTag(total int, tagId int, tagName string) StatisticWithTag {
	return &statisticWithTag{tagId: tagId, total: total, tagName: tagName}
}

func (s *statisticWithTag) Total() int {
	return s.total
}

func (s *statisticWithTag) TagId() int {
	return s.tagId
}

func (s *statisticWithTag) TagName() string {
	return s.tagName
}

func (s *statisticWithTag) ToResponse() {

}
