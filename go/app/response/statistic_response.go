package response

type TotalByArchiveType struct {
	Total int `json:"total"`
}

type TagStatistic struct {
	Total   int    `json:"total"`
	TagId   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
}

type TagStatisticListResponse struct {
	Statistics []TagStatistic `json:"statistics"`
}
