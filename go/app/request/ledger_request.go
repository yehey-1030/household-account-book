package request

import "time"

type LedgerListRequest struct {
	ArchiveTypeId int       `json:"archive_type_id"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
}
