package request

type UriArchiveTypeId int

type StatisticDateRangeRequest struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}
