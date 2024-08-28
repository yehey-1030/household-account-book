package response

import "time"

type LedgerResponse struct {
	LedgerId      int       `json:"ledgerID"`
	Title         string    `json:"title"`
	Memo          string    `json:"memo"`
	Amount        int       `json:"amount"`
	Date          time.Time `json:"date"`
	IsExcluded    bool      `json:"isExcluded"`
	ArchiveTypeId int       `json:"archiveTypeId"`
}

type LedgerListResponse struct {
	LedgerList []LedgerResponse `json:"ledger_list"`
}
