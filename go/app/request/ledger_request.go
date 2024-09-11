package request

type LedgerListRequest struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}

type LedgerCreateRequest struct {
	ArchiveTypeId int    `json:"archive_type_id"`
	Title         string `json:"title"`
	Memo          string `json:"memo"`
	Amount        int    `json:"amount"`
	Date          string `json:"date"`
	IsExcluded    bool   `json:"is_excluded"`
	Tags          []int  `json:"tags"`
}
