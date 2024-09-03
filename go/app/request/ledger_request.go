package request

import "time"

type LedgerListRequest struct {
	ArchiveTypeId int       `json:"archive_type_id"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
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
