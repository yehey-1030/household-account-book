package response

import "time"

type LedgerResponse struct {
	Title  string    `json:"title"`
	Memo   string    `json:"memo"`
	Amount int       `json:"amount"`
	Date   time.Time `json:"date"`
}

type LedgerListResponse struct {
	LedgerList []LedgerResponse `json:"ledger_list"`
}
