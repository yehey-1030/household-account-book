package response

import "time"

type LedgerResponse struct {
	LedgerId      int                 `json:"ledgerID"`
	Title         string              `json:"title"`
	Memo          string              `json:"memo"`
	Amount        int                 `json:"amount"`
	Date          time.Time           `json:"date"`
	IsExcluded    bool                `json:"isExcluded"`
	ArchiveTypeId ArchiveTypeResponse `json:"archiveTypeId"`
	Tags          []TagResponse       `json:"tags"`
}

type LedgerListResponse struct {
	LedgerList []LedgerResponse `json:"ledger_list"`
}
