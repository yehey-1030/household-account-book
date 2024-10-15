package domain

import "github.com/yehey-1030/household-account-book/go/app/response"

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
	ToResponse() response.TagStatistic
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

func (s *statisticWithTag) ToResponse() response.TagStatistic {
	return response.TagStatistic{TagId: s.tagId, TagName: s.tagName, Total: s.total}
}
